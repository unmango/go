package iter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Seq", func() {
	Context("Empty", func() {
		It("should not yield any elements", func() {
			seq := iter.Empty[int]()

			Expect(slices.Collect(seq)).To(BeEmpty())
		})
	})

	Context("Singleton", func() {
		It("should yield a single element", func() {
			seq := iter.Singleton(69)

			Expect(slices.Collect(seq)).To(ConsistOf(69))
		})
	})
})
