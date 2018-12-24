SUBDIRS = ./config ./scanner ./sh ./util ./rpc

main:
	env GOOS=linux GOARCH=amd64 go build -o main

clean:
	rm -f main main.zip

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

start-server: deps generate
	cd cmd/server \
		&& go build -o server \
		&& ./server

run-client: deps generate
	cd cmd/client \
		&& go build -o client \
		&& ./client

build-binaries: deps generate
	cd cmd/client && go build -o client
	cd cmd/server && go build -o server
