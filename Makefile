PKG_NAME=git-confed
DUSER=aceberg

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/$(PKG_NAME) && \
	go mod tidy

run:
	cd cmd/$(PKG_NAME)/ && \
	go run . -c /data/$(PKG_NAME)/config.yaml #-r /data/$(PKG_NAME)/repos.yaml 
	
fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/$(PKG_NAME)/ && \
	CGO_ENABLED=0 go build -o ../../tmp/$(PKG_NAME) .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .