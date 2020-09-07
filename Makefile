SHELL=/bin/bash
MAKEFILE_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_DIR))))
DIR=$(MAKEFILE_DIR)

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
# current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
current_dir := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: deps install clean
export GO111MODULE=on
export GOPATH=/go

# target:=$(GOPATH)/bin/$(notdir $(PWD))
target:=$(current_dir)/bin/applytmpl

all: $(target)

# all:
# 	echo $(mkfile_path)
# 	echo $(current_dir)
# 	echo $(wildcard *.go cmd/applytmpl/*.go)
# 	@echo MAKEFILE_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
# 	@echo current_dir := $(current_dir)
# 	@echo CURRENT_DIR := $(notdir $(MAKEFILE_DIR))
# 	@echo CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_DIR))))
# 	@echo DIR=$(MAKEFILE_DIR)


$(target): Makefile $(wildcard *.go cmd/applytmpl/*.go)
# 	# CGO_ENABLED=0 go install -tags netgo

# export HOSTNAME=$(shell hostname)
export dirty=$$(git diff --no-ext-diff --quiet|| echo \-dirty)
# $(target) : $(wildcard *.go) Makefile
export args="-s -w -X main.Version=$(shell tag=$$(git tag --points-at HEAD); if [[ -z $${tag:-} ]]; then echo untagged-commit; else echo $${tag}; fi) -X main.Build=$$(date -u +%Y.%m.%d.%H.%M.%S.%z) -X main.Commit=$$(git log --format=%h${dirty}-%aI -n1)"

.PHONY: build
build: $(target)


.dep: $(target) Makefile $(wildcard $(current_dir)/*.go $(current_dir)/cmd/applytmpl/*.go)
	touch .dep

%: $(current_dir)/bin/%

$(current_dir)/bin/%: $(dep)
	@echo "Building via % rule for $@ from $<"
	mkdir -p bin
	go build --tags netgo -ldflags=${args} -o $(current_dir)/bin/ $(current_dir)/cmd/...;

# cd cmd/applytmpl; CGO_ENABLED=0 go build --tags netgo -ldflags "$${args}" -o $@ $^ ;

install: build
	cp $(target) /go/bin/

clean:
	@if [[ -x "$(target)" ]]; then rm -f $(target); fi
	# @if [[ -d "bin" ]]; then rmdir bin; fi
.PHONY: test
test:
	make -C test
