# This Makefile is meant to be used by all-in-one build of aergo project

.PHONY: build all test clean libtool libtool-clean deps protoc protoclean

BINPATH := $(shell pwd)/bin
CMDS := aergocli aergosvr aergoluac
REPOPATH := github.com/aergoio/aergo

build: vendor libtool
	GOBIN=$(BINPATH) \
CGO_CFLAGS="-I$(LIBPATH)/include" \
CGO_LDFLAGS="-L$(LIBPATH)/lib" \
go install ./cmd/...

all: clean test build
	@echo "Done All"

vendor: glide.yaml glide.lock
	@glide install

# test

test:
	@CGO_CFLAGS="-I$(LIBPATH)/include" \
	CGO_LDFLAGS="-L$(LIBPATH)/lib" \
	go test -timeout 60s ./...

# clean

clean: libtool-clean
	go clean
	rm -f $(addprefix $(BINPATH)/,$(CMDS))

# 3rd party libs

LIBPATH := $(shell pwd)/libtool

libtool: 
	$(MAKE) -C $(LIBPATH) install

libtool-clean:
	$(MAKE) -C $(LIBPATH) uninstall

# etc

deps: vendor libtool
	@glide install

protoc:
	protoc -I/usr/local/include \
		-I${GOPATH}/src/${REPOPATH}/aergo-protobuf/proto \
		--go_out=plugins=grpc:${GOPATH}/src \
		${GOPATH}/src/${REPOPATH}/aergo-protobuf/proto/*.proto
	go build ./types/...

protoclean:
	rm -f types/*.pb.go

