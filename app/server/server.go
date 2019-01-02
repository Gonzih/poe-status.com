package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/Gonzih/poe-status.com/app/config"
	"gitlab.com/Gonzih/poe-status.com/db"
	"gitlab.com/Gonzih/poe-status.com/host"
	"gitlab.com/Gonzih/poe-status.com/migrations"
	"gitlab.com/Gonzih/poe-status.com/rpc"
	"gitlab.com/Gonzih/poe-status.com/web/assets"
	"gitlab.com/Gonzih/poe-status.com/web/endpoints"
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

	cfg, err := config.ReadYAML()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	inputToken := req.AuthToken.Token
	req.AuthToken = nil

	for name, token := range cfg.Tokens {
		if token == inputToken {
			log.Printf(`Got auth using "%s" token`, name)
			return true
		}
	}

	return false
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
