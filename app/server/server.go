package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Gonzih/poe-status.com/app/config"
	"github.com/Gonzih/poe-status.com/db"
	"github.com/Gonzih/poe-status.com/host"
	"github.com/Gonzih/poe-status.com/migrations"
	"github.com/Gonzih/poe-status.com/rpc"
	"github.com/Gonzih/poe-status.com/web/assets"
	"github.com/Gonzih/poe-status.com/web/endpoints"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
)

// Options repsenend command line options
type Options struct {
	Host                 string
	Port                 int
	DatabaseURL          string
	MigrationsFolderPath string
}

// checkToken will validate token for ScanResults and reset the field
func checkToken(req *rpc.ScanResults) bool {
	if req.AuthToken == nil {
		log.Printf("Got empty token!")
		return false
	}

	token := config.MainToken()
	if len(token) == 0 {
		log.Fatal("Could not load token")
		return false
	}

	inputToken := req.AuthToken.Token
	req.AuthToken = nil

	return token == inputToken
}

// PoeStatusServer implements Twirp server
type PoeStatusServer struct{}

// SaveScanResults implements Twirp handler
func (s *PoeStatusServer) SaveScanResults(ctx context.Context, req *rpc.ScanResults) (*rpc.Empty, error) {
	tokenValid := checkToken(req)

	if !tokenValid {
		return &rpc.Empty{}, errors.New("Got invalid token")
	}

	buf := bytes.NewBuffer(nil)
	marshaller := jsonpb.Marshaler{EmitDefaults: true}
	err := marshaller.Marshal(buf, req)
	if err != nil {
		return &rpc.Empty{}, err
	}

	ts, err := ptypes.Timestamp(req.CreatedAt)
	if err != nil {
		return &rpc.Empty{}, err
	}

	up, err := host.IsUp(req)
	if err != nil {
		log.Printf("Host IsUp error was: %s", err)
	}

	result := &db.ScanResult{
		ScanIP:    req.ScanIP,
		Host:      req.Host,
		Up:        up,
		CreatedAt: ts,
		QueryData: buf.Bytes(),
	}

	err = db.NewDefaultScanResultStore().SaveScanResult(result)
	if err != nil {
		return &rpc.Empty{}, err
	}

	return &rpc.Empty{}, nil
}

// StartServer will start http server
func StartServer(opts *Options) error {
	err := migrations.Up(opts.MigrationsFolderPath, opts.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Init(opts.DatabaseURL)
	if err != nil {
		return err
	}

	go func() {
		for {
			log.Println("Doing db cleanup")
			err := db.NewDefaultScanResultStore().DeleteResultsOlderThan(time.Hour * 5)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Minute * 30)
		}
	}()

	bindAddr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	log.Printf("Starting server on %s", bindAddr)
	twirpHandler := rpc.NewPoeStatusServer(&PoeStatusServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(safeFileSystem{fs: assets.AssetsDir}))
	mux.HandleFunc("/data.json", endpoints.ScanAggrEndpoint)
	mux.Handle(rpc.PoeStatusPathPrefix, twirpHandler)

	return http.ListenAndServe(bindAddr, mux)
}
