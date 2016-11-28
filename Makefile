CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-uri; then rm -rf src/github.com/whosonfirst/go-whosonfirst-uri; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-uri
	cp uri.go src/github.com/whosonfirst/go-whosonfirst-uri/uri.go

fmt:	self
	go fmt uri.go
	go fmt cmd/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-expand cmd/wof-expand.go
