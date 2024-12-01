_ := $(shell mkdir -p .make bin)

REPO := github.com/unmango/go
PKGS := iter maps result iter/seqs slices rx rx/observable option maybe

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVOPS := ${LOCALBIN}/devops
GINKGO := ${LOCALBIN}/ginkgo

ifeq ($(CI),)
TEST_FLAGS := --label-filter !E2E
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build: bin/devops .make/build
test: .make/test
tidy: go.sum

test_all:
	$(GINKGO) run -r ./

tidy: go.sum

clean:
	find . -name report.json -delete

go.sum: go.mod $(shell $(DEVOPS) list --go)
	go mod tidy

%_suite_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

bin/ginkgo: go.mod
	GOBIN=${LOCALBIN} go install github.com/onsi/ginkgo/v2/ginkgo

# This recursive dependency works 90% of the time, and when it doesn't its easy to fix
bin/devops: $(shell $(DEVOPS) list --go --exclude-tests)
	go build -o $@ cmd/devops/main.go

.envrc: hack/example.envrc
	cp $< $@

.make/build: $(shell $(DEVOPS) list --go --exclude-tests)
	go build ./...
	@touch $@

.make/test: $(shell $(DEVOPS) list --go) | bin/ginkgo
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
