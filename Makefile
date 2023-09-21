install:
	go mod download

test:
	go test -v --cover ./...

build-mac:
	GOOS=darwin GOARCH=amd64 go build $(ARGS)
build-linux:
	GOOS=linux GOARCH=amd64 go build $(ARGS)

build-all:
# the binary names match the pattern:
# bin/<name>-${JFROG_OPERATING_SYSTEM_FAMILY}-${JFROG_ARCHITECTURE}
# this will allow the task to easily execute the correct one based on the node
	GOOS=linux GOARCH=amd64 go build $(ARGS) -o bin/hello-Linux-x86_64
	GOOS=linux GOARCH=arm64 go build $(ARGS) -o bin/hello-Linux-ARM64
	GOOS=darwin GOARCH=amd64 go build $(ARGS) -o bin/hello-Darwin-x86_64
	GOOS=darwin GOARCH=arm64 go build $(ARGS) -o bin/hello-Darwin-ARM64
	GOOS=windows GOARCH=amd64 go build $(ARGS) -o bin/hello-Windows-x86_64
