# TODO: Build using https://goreleaser.com/
build:
	CGO_ENABLED=0 go build -v

build-linux:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -v

release-linux: build-linux
	tar cvzf pg-query-killer-linux-amd64.tar.gz pg-query-killer

# TODO: Lint using golangci-lint
