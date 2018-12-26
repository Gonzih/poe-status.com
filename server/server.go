package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/Gonzih/poe-status.com/db"
	"gitlab.com/Gonzih/poe-status.com/host"
	"gitlab.com/Gonzih/poe-status.com/migrations"
	"gitlab.com/Gonzih/poe-status.com/rpc"
	"gitlab.com/Gonzih/poe-status.com/web/ui"
)

// Options repsenend command line options
type Options struct {
	Host                 string
	Port                 int
	DatabaseURL          string
	MigrationsFolderPath string
}

type PoeStatusServer struct{}

func (s *PoeStatusServer) SaveScanResults(ctx context.Context, req *rpc.ScanResults) (*rpc.Empty, error) {
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

	err = db.SaveScanResult(db.DB(), result)
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

	bindAddr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	log.Printf("Starting server on %s", bindAddr)
	twirpHandler := rpc.NewPoeStatusServer(&PoeStatusServer{}, nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.Handle(rpc.PoeStatusPathPrefix, twirpHandler)

	return http.ListenAndServe(bindAddr, mux)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	err := ui.Index(res)
	if err != nil {
		log.Fatal(err)
	}
}
