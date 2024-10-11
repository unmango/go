REPO := github.com/unmango/go
PKGS := iter maps result iter/seqs slices rx rx/obs

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

GINKGO := ${LOCALBIN}/ginkgo

ifeq ($(CI),)
TEST_FLAGS := --json-report report.json --keep-separate-reports
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build:
	go build ./...

tidy: go.mod
	go mod tidy

test: $(addsuffix /report.json,${PKGS})

clean:
	find . -name report.json -delete

$(addsuffix /report.json,${PKGS}): %/report.json: $(GINKGO) $(wildcard %/*.go)
	$< run ${TEST_FLAGS} $*

%_suite_test.go: | $(GINKGO)
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go: | $(GINKGO)
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

$(LOCALBIN):
	mkdir -p $@

$(GINKGO): .versions/ginkgo $(LOCALBIN)
	GOBIN=${LOCALBIN} go install github.com/onsi/ginkgo/v2/ginkgo@v$(shell cat $<)
