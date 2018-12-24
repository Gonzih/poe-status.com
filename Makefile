SUBDIRS = ./config ./scanner ./sh ./util

main:
	env GOOS=linux GOARCH=amd64 go build -o main

main.zip: main
	zip main.zip main servers.yaml
	ls -lha main.zip


clean:
	rm -f main main.zip

run: main
	./main

test:
	go test -v --cover $(SUBDIRS)

generate:
	go generate $(SUBDIRS)
