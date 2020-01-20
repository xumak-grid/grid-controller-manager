.PHONY: run test build compile image deploy

build: 
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/grid-controller  -tags netgo cmd/grid-controller/*.go
