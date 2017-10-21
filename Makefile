target:=$(GOPATH)/bin/$(notdir $(PWD))

all: $(target)

$(target): Makefile $(wildcard *.go)
	CGO_ENABLED=0 go install -tags netgo
