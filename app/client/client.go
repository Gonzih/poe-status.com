package client

import (
	"context"
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
		var pingInfo *rpc.PingInfo

		if host.Platform == "PC" {
			if nmapAvailable {
				ports, err = scanner.NmapScan(host.Host)
			} else {
				ports, err = scanner.GoScan(host.Host, cfg.Ports)
			}

			if err != nil {
				scanError = err.Error()
			}
		} else {
			pingInfo, err = scanner.Ping(host.Host, 100)
			if err != nil {
				scanError = err.Error()
			}
		}

		_, err := client.SaveScanResults(context.Background(), &rpc.ScanResults{
			ScanIP:    myip.String(),
			Host:      host.Host,
			CreatedAt: ptypes.TimestampNow(),
			Ports:     ports,
			ScanError: scanError,
			Platform:  host.Platform,
			PingInfo:  pingInfo,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
