package iter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Seq3", func() {
	Describe("Empty", func() {
		It("should not yield any elements", func() {
			seq := iter.Empty3[int, string, bool]()

			a, b, c := slices.Collect3(seq)
			Expect(a).To(BeEmpty())
			Expect(b).To(BeEmpty())
			Expect(c).To(BeEmpty())
		})
	})

	Describe("Singleton", func() {
		It("should yield a single element", func() {
			seq := iter.Singleton3(69, "420", true)

			a, b, c := slices.Collect3(seq)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("420"))
			Expect(c).To(ConsistOf(true))
		})
	})

	Describe("DropFirst", func() {
		It("should yield the last two elements of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.DropFirst3(seq)

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf("420"))
			Expect(b).To(ConsistOf(true))
		})
	})

	Describe("DropMid", func() {
		It("should yield the first and last elements of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.DropMid3(seq)

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf(true))
		})
	})

	Describe("DropLast", func() {
		It("should yield the first two elements of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.DropLast3(seq)

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("420"))
		})
	})

	Describe("KeepFirst", func() {
		It("should yield the first element of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.KeepFirst3(seq)

			a := slices.Collect(r)
			Expect(a).To(ConsistOf(69))
		})
	})

	Describe("KeepMid", func() {
		It("should yield the middle element of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.KeepMid3(seq)

			a := slices.Collect(r)
			Expect(a).To(ConsistOf("420"))
		})
	})

	Describe("KeepLast", func() {
		It("should yield the last element of the triple", func() {
			seq := iter.Singleton3(69, "420", true)

			r := iter.KeepLast3(seq)

			a := slices.Collect(r)
			Expect(a).To(ConsistOf(true))
		})
	})
})
