# Go parameters
#export GO111MODULE=on
#export GOPATH:=${GOPATH}:$(shell pwd)/../../../
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=example

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) example.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
