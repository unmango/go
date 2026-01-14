_ := $(shell mkdir -p .make bin)

REPO := github.com/unmango/go
PKGS := iter maps result iter/seqs slices rx rx/observable option maybe

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

export GOBIN := ${LOCALBIN}

GO        ?= go
DEVCTL    ?= $(GO) tool devctl
GINKGO    ?= $(GO) tool ginkgo
GOMOD2NIX ?= $(GO) tool gomod2nix
NIX       ?= nix

ifeq ($(CI),)
TEST_FLAGS := --label-filter !E2E
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build: .make/build bin/nix
test: .make/test
tidy: go.sum

test_all:
	$(GINKGO) run -r ./

clean:
	find . -name report.json -delete

bin/nix:
	$(NIX) build --out-link $@

go.sum: go.mod $(shell $(DEVCTL) list --go)
	go mod tidy

gomod2nix.toml: go.mod
	$(GOMOD2NIX)

%_suite_test.go:
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go:
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

.envrc: hack/example.envrc
	cp $< $@

.make/build: $(shell $(DEVCTL) list --go --exclude-tests)
	go build ./...
	@touch $@

.make/nix-build:
	$(NIX) build

.make/test: $(shell $(DEVCTL) list --go)
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
