.PHONY: all

all: run

run: cmd/main/main.go
	@go run $<

resolve: scripts/resolve.sh
	@sh $<

build: cmd/main/main.go
	@go build $<