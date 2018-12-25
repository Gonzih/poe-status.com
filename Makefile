.PHONY=generate test run clean deps start-server run-client build-binaries clean

BINARY_DIRS = cmd/poe-status-client cmd/poe-status-server
SUBDIRS = ./config ./scanner ./sh ./util ./rpc ./myip ./db ./migrations ./host
GENERATE_FILES = rpc/service.pb.go rpc/service.twirp.go config/datadir_vfsdata.go

test: generate
	go test -v --cover $(SUBDIRS)

generate: deps
	go generate $(SUBDIRS)

deps:
	go get
	go get github.com/twitchtv/twirp/protoc-gen-twirp
	go get github.com/golang/protobuf/protoc-gen-go

start-server: generate
	cd cmd/poe-status-server && go build -o poe-status-server && ./poe-status-server --migrations-folder=../../migrations --database=postgres://postgres@localhost/poetest?sslmode=disable

run-client: generate
	cd cmd/poe-status-client && go build -o poe-status-client && ./poe-status-client

build-binaries: generate
	cd cmd/poe-status-client && go build -o poe-status-client
	cd cmd/poe-status-server && go build -o poe-status-server

clean:
	rm -f $(GENERATE_FILES) main

$(GENERATE_FILES): generate
