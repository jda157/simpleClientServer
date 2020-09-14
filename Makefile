.PHONY: build
build:
	go build -o bin/client client/client.go
	go build -o bin/server server/server.go

.PHONY: clean
clean:
	rm -rf bin
	
