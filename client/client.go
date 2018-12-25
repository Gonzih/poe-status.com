package client

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/Gonzih/poe-status.com/rpc"
)

// Options repsenend command line options
type Options struct {
	URL string
}

// Call is the cli entry point
func Call(opts *Options) error {
	client := rpc.NewPoeStatusProtobufClient(opts.URL, &http.Client{})

	resp, err := client.SaveScanResults(context.Background(), &rpc.ScanResults{})
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}