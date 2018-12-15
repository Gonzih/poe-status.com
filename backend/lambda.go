package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

const nmapPath = "./usr/bin/nmap"

func HandleRequest(ctx context.Context) ([]*Host, error) {
	data := readServersYaml()

	hosts := make([]*Host, 0)
	for name, host := range data.Pc {
		h := &Host{Name: name, HostName: host}
		log.Println(name)
		up, err := h.IsUp()
		log.Printf("%s is %v: %v", name, up, err)
		hosts = append(hosts, h)
	}

	return hosts, nil
}

func lambdaMain() {
	lambda.Start(HandleRequest)
}
