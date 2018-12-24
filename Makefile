.PHONY=generate test run clean deps start-server run-client build-binaries clean

SUBDIRS = ./config ./scanner ./sh ./util ./rpc
GENERATE_FILES = rpc/service.pb.go rpc/service.twirp.go config/datadir_vfsdata.go

main:
	env GOOS=linux GOARCH=amd64 go build -o main

run: main
	./main

test: generate
	go test -v --cover $(SUBDIRS)

generate: deps
	go generate $(SUBDIRS)

deps:
	go get
	go get github.com/twitchtv/twirp/protoc-gen-twirp
	go get github.com/golang/protobuf/protoc-gen-go

start-server: generate
	cd cmd/server \
		&& go build -o server \
		&& ./server

run-client: generate
	cd cmd/client \
		&& go build -o client \
		&& ./client

build-binaries: generate
	cd cmd/client && go build -o client
	cd cmd/server && go build -o server

clean:
	rm -f main main.zip
	rm -f $(GENERATE_FILES) main

$(GENERATE_FILES): generate
