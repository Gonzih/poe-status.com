package client

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"

	ptypes "github.com/golang/protobuf/ptypes"
	"gitlab.com/Gonzih/poe-status.com/app/config"
	"gitlab.com/Gonzih/poe-status.com/myip"
	"gitlab.com/Gonzih/poe-status.com/rpc"
	"gitlab.com/Gonzih/poe-status.com/scanner"
	"golang.org/x/sync/semaphore"
)

// Options repsenend command line options
type Options struct {
	URL string
}

var pingCount = 50

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
	pingAvailable := scanner.PingAvailable()

	if !nmapAvailable {
		log.Fatal("Please install nmap and make sure it's available on your PATH")
	}

	if !pingAvailable {
		log.Fatal("Please install ping and make sure it's available on your PATH")
	}

	var wg sync.WaitGroup
	nmapLock := semaphore.NewWeighted(3)
	errChan := make(chan error, 500)

	for _, host := range cfg.AllHosts() {
		wg.Add(1)

		go func(host config.Host) {
			defer wg.Done()

			var ports []*rpc.PortInfo
			var err error
			var scanError string
			var pingInfo *rpc.PingInfo

			if host.Platform == "PC" {
				if nmapAvailable {
					nmapLock.Acquire(context.TODO(), 1)
					ports, err = scanner.NmapScan(host.Host)
					nmapLock.Release(1)
				}

				if err != nil {
					scanError = err.Error()
				}
			}

			pingInfo, err = scanner.Ping(host.Host, pingCount)
			if err != nil {
				scanError = err.Error()
			}

			_, err = client.SaveScanResults(context.Background(), &rpc.ScanResults{
				ScanIP:    myip.String(),
				Host:      host.Host,
				CreatedAt: ptypes.TimestampNow(),
				Ports:     ports,
				ScanError: scanError,
				Platform:  host.Platform,
				PingInfo:  pingInfo,
				AuthToken: &rpc.AuthToken{
					Token: cfg.Tokens["main"],
				},
			})

			if err != nil {
				errChan <- err
			}
		}(host)
	}

	wg.Wait()

	close(errChan)

	var errStrings []string
	for err := range errChan {
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}
	if len(errStrings) > 0 {
		return errors.New(strings.Join(errStrings, " "))
	}

	return nil
}
