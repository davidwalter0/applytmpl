.PHONY: deps install clean
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

# $(target) : $(wildcard *.go) Makefile

build: $(target)


.dep: $(target) Makefile
	touch .dep

%: bin/%

bin/%: %.go 
	@echo "Building via % rule for $@ from $<"
	@if go version|grep -q 1.4 ; then											\
	    args="-s -w -X main.Build $$(date -u +%Y.%m.%d.%H.%M.%S.%:::z) -X main.Commit $$(git log --format=%hash-%aI -n1)";	\
	fi;															\
	if go version|grep -qE "(1\.[5-9](\.?[0-9])*|1.[1-9][0-9](\.?[0-9])+|2.[0-9](\.?[0-9])*)"; then				\
	    args="-s -w -X main.Build=$$(date -u +%Y.%m.%d.%H.%M.%S.%:::z) -X main.Commit=$$(git log --format=%hash-%aI -n1)";	\
	fi;															\
	CGO_ENABLED=0 go build --tags netgo -ldflags "$${args}" -o $@ $^ ;

install: build
	cp $(target) /go/bin/

clean:
	@if [[ -x "$(target)" ]]; then rm -f $(target); fi
	@if [[ -d "bin" ]]; then rmdir bin; fi

govendor:
	govendor init
	govendor add +external

govendor-sync:
	govendor sync
