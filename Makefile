# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=BiliDownloader

# Determine the OS and architecture
ifeq ($(OS),Windows_NT)
    OSFLAG := windows
    EXTENSION := .exe
else
    OSFLAG := $(shell uname | tr '[:upper:]' '[:lower:]')
    EXTENSION :=
endif

build: $(BINARY_NAME)

$(BINARY_NAME):
	$(GOBUILD) -o $(BINARY_NAME)$(EXTENSION) -v  ./cmd

# Build for Linux
linux:
	go env -w CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"
	$(GOBUILD) -o $(BINARY_NAME)_linux_amd64 -v ./cmd
	go env -w CGO_ENABLED="1" GOOS="windows" GOARCH="amd64"

# Build for Mac
darwin_x86:
	go env -w CGO_ENABLED="0" GOOS="darwin" GOARCH="amd64"
	$(GOBUILD) -o $(BINARY_NAME)_darwin_amd64 -v ./cmd
	go env -w CGO_ENABLED="1" GOOS="windows" GOARCH="amd64"

darwin_arm:
	go env -w CGO_ENABLED="0" GOOS="darwin" GOARCH="arm64"
	$(GOBUILD) -o $(BINARY_NAME)_darwin_arm64 -v ./cmd
	go env -w CGO_ENABLED="1" GOOS="windows" GOARCH="amd64"

# Build for Windows
windows:
	go env -w CGO_ENABLED="1" GOOS="windows" GOARCH="amd64"
	$(GOBUILD) -o $(BINARY_NAME)_windows_amd64$(EXTENSION) -v ./cmd

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)*$(EXTENSION)