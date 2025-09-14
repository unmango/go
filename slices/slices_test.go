package slices_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Slices", func() {
	It("Should compact a seq", func() {
		seq := iter.Append(iter.Empty[int](), 1, 2, 2, 3)

		res := slices.CompactSeq(seq)

		Expect(res).To(Equal([]int{1, 2, 3}))
	})

	It("Should compact a seq with an equality function", func() {
		seq := iter.Append(iter.Empty[int](), 1, 2, 2, 3)

		res := slices.CompactSeqFunc(seq, func(a, b int) bool {
			return a == b
		})

		Expect(res).To(Equal([]int{1, 2, 3}))
	})
})
