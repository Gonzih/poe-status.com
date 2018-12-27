package client

import (
	"context"
	"log"
	"net/http"

	ptypes "github.com/golang/protobuf/ptypes"
	"gitlab.com/Gonzih/poe-status.com/app/config"
	"gitlab.com/Gonzih/poe-status.com/myip"
	"gitlab.com/Gonzih/poe-status.com/rpc"
	"gitlab.com/Gonzih/poe-status.com/scanner"
)

// Options repsenend command line options
type Options struct {
	URL string
}

// Call is the cli entry point
func Call(opts *Options) error {
	cfg, err := config.ReadYAML()
	if err != nil {
		return err
	}

	myip, err := myip.MyIP()
	if err != nil {
		return err
	}

	client := rpc.NewPoeStatusProtobufClient(opts.URL, &http.Client{})
	nmapAvailable := scanner.NmapAvailable()

	for _, host := range cfg.AllHosts() {
		var ports []*rpc.PortInfo
		var err error
		var scanError string

		if nmapAvailable {
			ports, err = scanner.NmapScan(host.Host)
		} else {
			ports, err = scanner.GoScan(host.Host, cfg.Ports)
		}

		if err != nil {
			scanError = err.Error()
		}

		resp, err := client.SaveScanResults(context.Background(), &rpc.ScanResults{
			ScanIP:    myip.String(),
			Host:      host.Host,
			CreatedAt: ptypes.TimestampNow(),
			Ports:     ports,
			ScanError: scanError,
		})
		if err != nil {
			return err
		}
		log.Println(resp)
	}

	return nil
}
