.PHONY: local-db clean-local-env unit-test run

local-db:
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"

clean-local-env:
	eval "docker-compose -f localdb-docker-compose.yaml down"

unit-test:
	go test ./... -tags=unit -count=1

build: 
	go build -o server *.go

run: build
	ENV=local ./server; rm server