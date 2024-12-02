package iter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Seq2", func() {
	Describe("Empty", func() {
		It("should not yield any elements", func() {
			seq := iter.Empty2[int, string]()

			a, b := slices.Collect2(seq)
			Expect(a).To(BeEmpty())
			Expect(b).To(BeEmpty())
		})
	})

	Describe("Singleton", func() {
		It("should yield a single element", func() {
			seq := iter.Singleton2(69, "420")

			a, b := slices.Collect2(seq)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("420"))
		})
	})

	Describe("DropFirst", func() {
		It("should yield the last element of the tuple", func() {
			seq := iter.Singleton2(69, "420")

			r := iter.DropFirst2(seq)

			Expect(slices.Collect(r)).To(ConsistOf("420"))
		})
	})

	Describe("DropLast", func() {
		It("should yield the first element of the tuple", func() {
			seq := iter.Singleton2(69, "420")

			r := iter.DropLast2(seq)

			Expect(slices.Collect(r)).To(ConsistOf(69))
		})
	})
})
