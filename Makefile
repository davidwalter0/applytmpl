.PHONY: deps install clean
export GO111MODULE=on
export GOPATH=/go

# target:=$(GOPATH)/bin/$(notdir $(PWD))
target:=bin/applytmpl

all: $(target)

$(target): $(wildcard *.go)
# 	# CGO_ENABLED=0 go install -tags netgo

SHELL=/bin/bash
MAKEFILE_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_DIR))))
DIR=$(MAKEFILE_DIR)
# export HOSTNAME=$(shell hostname)
export dirty:=$(shell git diff --no-ext-diff --quiet|| echo \-dirty)
# $(target) : $(wildcard *.go) Makefile
.PHONY: build
build: $(target)


.dep: $(target) Makefile
	touch .dep

%: bin/%

bin/%: %.go 
	@echo "Building via % rule for $@ from $<"
	if go version|grep -qE "(1\.[5-9](\.?[0-9])*|1.[1-9][0-9](\.?[0-9])+|2.[0-9](\.?[0-9])*)"; then				\
	    args="-s -w -X main.Build=$$(date -u +%Y.%m.%d.%H.%M.%S.%:::z) -X main.Commit=$$(git log --format=%hash$${dirty}-%aI -n1)";	\
	fi;															\
	CGO_ENABLED=0 go build --tags netgo -ldflags "$${args}" -o $@ $^ ;

install: build
	cp $(target) /go/bin/

clean:
	@if [[ -x "$(target)" ]]; then rm -f $(target); fi
	@if [[ -d "bin" ]]; then rmdir bin; fi
.PHONY: test
test:
	make -C test
