BINARY_NAME=rgen
BINARY_NAME_LINUX=$(BINARY_NAME)-linux
BINARY_NAME_DARWIN=$(BINARY_NAME)
BINARY_NAME_WINDOWS=$(BINARY_NAME)-windows

export GO111MODULE=on

all: check lint build

build:
	go build -o $(BINARY_NAME) -v ./
	chmod +x $(BINARY_NAME)
test:
	go test ./...
lint:
	golint ./
clean:
	go clean
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_NAME_LINUX)
	rm -rf $(BINARY_NAME_DARWIN)
	rm -rf $(BINARY_NAME_WINDOWS)
	

build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/$(BINARY_NAME_LINUX) -v ./

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/$(BINARY_NAME_DARWIN) -v ./
	
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/$(BINARY_NAME_WINDOWS) -v ./
