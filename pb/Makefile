# Go parameters
#export GO111MODULE=off
#export GOPATH:=$(shell pwd)/../
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=example
TEST_NAME=test
all: build test
build:
	$(GOBUILD) -o $(BINARY_NAME) example.go
test:
	$(GOBUILD) -o $(TEST_NAME) test.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(TEST_NAME)
