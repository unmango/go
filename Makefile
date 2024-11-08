_ := $(shell mkdir -p .make bin)

REPO := github.com/unmango/go
PKGS := iter maps result iter/seqs slices rx rx/observable option

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVOPS := ${LOCALBIN}/devops
GINKGO := ${LOCALBIN}/ginkgo

ifeq ($(CI),)
TEST_FLAGS := --json-report report.json --keep-separate-reports
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build: bin/devops .make/build
test: .make/test

tidy: go.mod
	go mod tidy

clean:
	find . -name report.json -delete

%_suite_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

bin/ginkgo: go.mod
	GOBIN=${LOCALBIN} go install github.com/onsi/ginkgo/v2/ginkgo

# This recursive dependency works 90% of the time, and when it doesn't its easy to fix
bin/devops: $(shell $(DEVOPS) list --go --exclude-tests)
	go build -o $@ cmd/devops/main.go

.make/build: $(shell $(DEVOPS) list --go --exclude-tests)
	go build ./...
	@touch $@

.make/test: $(shell $(DEVOPS) list --go) | bin/ginkgo
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
