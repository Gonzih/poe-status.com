package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/Gonzih/poe-status.com/rpc"
)

type PoeStatusServer struct{}

func (s *PoeStatusServer) SaveScanResults(ctx context.Context, req *rpc.ScanResults) (*rpc.Empty, error) {
	log.Printf("%#v", req)
	return &rpc.Empty{}, nil
}

func StartServer(port int) error {
	twirpHandler := rpc.NewPoeStatusServer(&PoeStatusServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(rpc.PoeStatusPathPrefix, twirpHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
