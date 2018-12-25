package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/Gonzih/poe-status.com/rpc"
)

// Options repsenend command line options
type Options struct {
	Host        string
	Port        int
	DatabaseURL string
}

type PoeStatusServer struct{}

func (s *PoeStatusServer) SaveScanResults(ctx context.Context, req *rpc.ScanResults) (*rpc.Empty, error) {
	log.Printf("%#v", req)
	return &rpc.Empty{}, nil
}

func StartServer(opts *Options) error {
	bindAddr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	log.Printf("Starting server on %s", bindAddr)
	twirpHandler := rpc.NewPoeStatusServer(&PoeStatusServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(rpc.PoeStatusPathPrefix, twirpHandler)
	return http.ListenAndServe(bindAddr, mux)
}
