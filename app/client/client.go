package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Gonzih/poe-status.com/app/config"
	"github.com/Gonzih/poe-status.com/myip"
	"github.com/Gonzih/poe-status.com/rpc"
	"github.com/Gonzih/poe-status.com/scanner"
	ptypes "github.com/golang/protobuf/ptypes"
	"golang.org/x/sync/semaphore"
)

// Options repsenend command line options
type Options struct {
	URL string
}

var pingCount = 50

// Call is the cli entry point
func Call(opts *Options) error {
	start := time.Now()
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
	pingLock := semaphore.NewWeighted(10)
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
					log.Printf("Nmap error for %s was %s", host.Host, err)
					scanError = err.Error()
				}

				open := 0
				for _, port := range ports {
					if port.Open {
						open++
					}
				}
				log.Printf("Host %s got %d ports from scan, %d open ports", host.Host, len(ports), open)
			}

			pingLock.Acquire(context.TODO(), 1)
			pingInfo, err = scanner.Ping(host.Host, pingCount)
			pingLock.Release(1)

			if err != nil {
				log.Printf("Ping error for %s was %s", host.Host, err)
				scanError = fmt.Sprintf("%s %s", scanError, err.Error())
			} else {
				log.Printf("Host %s got ping loss %d%%", host.Host, pingInfo.Loss)
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
					Token: config.MainToken(),
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

	log.Printf("Finished in %d seconds", int(time.Since(start).Seconds()))
	return nil
}
