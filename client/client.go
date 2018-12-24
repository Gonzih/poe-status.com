package client

import (
	"context"
	"fmt"
	"net/http"

	"gitlab.com/Gonzih/poe-status.com/rpc"
	"gitlab.com/Gonzih/poe-status.com/util"
)

func Call() {
	client := rpc.NewPoeStatusProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.SaveScanResults(context.Background(), &rpc.ScanResults{})
	util.Must(err)
	fmt.Println(resp)
}
