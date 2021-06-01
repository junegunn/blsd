ROOTDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOPATH  := $(ROOTDIR)/gopath
GO111MODULE := auto
export GOPATH GO111MODULE

blsd: git2go blsd.go
	go build -ldflags -w -o $@

$(GOPATH):
	mkdir -p $@

git2go: $(GOPATH)
	go get -d
	cd $(GOPATH)/src/github.com/libgit2/git2go && git checkout next && git submodule update --init && make install

run:
	go run blsd.go

clean:
	rm -rf $(GOPATH) blsd

.PHONY: run git2go clean
