REPO := github.com/unmango/go
MODS := iter maps result seqs slices

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

GINKGO := ${LOCALBIN}/ginkgo

ifeq ($(CI),)
TEST_FLAGS := --json-report report.json --keep-separate-reports
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build: $(MODS)
test: $(addsuffix /report.json,${MODS})
tidy: $(addsuffix /go.sum,${MODS})

clean:
	find . -name report.json -delete

.PHONY: ${MODS}
$(MODS):
	go -C $@ build ./...

$(addsuffix /report.json,${MODS}): %/report.json: $(GINKGO) $(wildcard %/*.go)
	$< run ${TEST_FLAGS} $*

$(addsuffix /go.mod,${MODS}): %/go.mod:
	mkdir -p $* && go -C $* mod init ${REPO}/$*

$(addsuffix /go.sum,${MODS}): %/go.sum: %/go.mod
	go -C $* mod tidy

%_suite_test.go: | $(GINKGO)
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go: | $(GINKGO)
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

go.work: | $(addsuffix /go.mod,${MODS})
	go work init ${MODS}

go.work.sum: go.work $(addsuffix /go.mod,${MODS})
	go work sync

$(LOCALBIN):
	mkdir -p $@

$(GINKGO): .versions/ginkgo $(LOCALBIN)
	GOBIN=${LOCALBIN} go install github.com/onsi/ginkgo/v2/ginkgo@v$(shell cat $<)
