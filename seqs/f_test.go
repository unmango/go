package seqs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/seqs"
	"github.com/unmango/go/slices"
)

var _ = Describe("F", func() {
	Context("Append", func() {
		It("should append an item to an empty sequence", func() {
			result := seqs.Append(iter.Empty[int](), 69)

			Expect(slices.Collect(result)).To(HaveExactElements(69))
		})

		It("should append an item to a non-empty sequence", func() {
			result := seqs.Append(iter.Singleton(69), 420)

			Expect(slices.Collect(result)).To(HaveExactElements(69, 420))
		})
	})
})
