# All this crap for working with modules in case I decide to switch back to it

MODS := iter maps result seqs slices

build: $(MODS)
test: $(addsuffix /report.json,${MODS})
tidy: $(addsuffix /go.sum,${MODS})

.PHONY: ${MODS}
$(MODS):
	go -C $@ build ./...

$(addsuffix /report.json,${MODS}): %/report.json: $(GINKGO) $(wildcard %/*.go)
	$< run ${TEST_FLAGS} $*

$(addsuffix /go.mod,${MODS}): %/go.mod:
	mkdir -p $* && go -C $* mod init ${REPO}/$*

$(addsuffix /go.sum,${MODS}): %/go.sum: %/go.mod $(wildcard %/*.go)
	go -C $* mod tidy

go.work: | $(addsuffix /go.mod,${MODS})
	go work init ${MODS}

go.work.sum: go.work $(addsuffix /go.mod,${MODS})
	go work sync
