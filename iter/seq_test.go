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
			v := iter.Empty[int]()

			Expect(slices.Collect(v)).To(BeEmpty())
		})
	})
})
