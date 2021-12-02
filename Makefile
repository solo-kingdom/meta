.PHONY: all

all: run

run: cmd/main/meta.go
	@go run $<

resolve: scripts/resolve.sh
	@sh $<

build: cmd/main/meta.go
	@go build $<