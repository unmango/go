package iter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

type bindStub struct {
	seq iter.Seq[int]
}

var _ = Describe("Seq", func() {
	Describe("Bind", func() {
		It("should bind", func() {
			a := iter.Singleton(69)
			b := iter.Singleton(bindStub{a})

			r := iter.Bind(b, func(s bindStub) iter.Seq[int] {
				return s.seq
			})

			Expect(r).To(ConsistOf(69))
		})
	})

	Describe("Empty", func() {
		It("should not yield any elements", func() {
			seq := iter.Empty[int]()

			Expect(seq).To(BeEmpty())
		})
	})

	Describe("Filter", func() {
		It("should not yield the filtered element", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Filter(s, func(i int) bool {
				return i == 69
			})

			Expect(r).To(ConsistOf(69))
		})
	})

	Describe("Head", func() {
		It("should yield the first element", func() {
			s := slices.Values([]int{69, 420})

			r, err := iter.Head(s)

			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal(69))
		})

		It("should error when the sequence is empty", func() {
			s := iter.Empty[int]()

			_, err := iter.Head(s)

			Expect(err).To(MatchError("empty sequence"))
		})
	})

	Describe("Flat", func() {
		It("should flatten a nested sequence", func() {
			a := iter.Singleton(69)
			b := iter.Singleton(420)
			c := slices.Values([]iter.Seq[int]{a, b})

			r := iter.Flat(c)

			Expect(r).To(ConsistOf(69, 420))
		})
	})

	Describe("FlatMap", func() {
		It("should flatten and map a nested sequence", func() {
			a := iter.Singleton(69)
			b := iter.Singleton(420)
			c := slices.Values([]iter.Seq[int]{a, b})

			r := iter.FlatMap(c, func(i int) int {
				return i + 1
			})

			Expect(r).To(ConsistOf(70, 421))
		})
	})

	Describe("Fold", func() {
		It("should fold", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Fold(s, func(sum int, x int) int {
				return sum + x
			}, 0)

			Expect(r).To(Equal(489))
		})
	})

	Describe("Fold", func() {
		It("should map", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Map(s, func(x int) int {
				return x + 1
			})

			Expect(r).To(ConsistOf(70, 421))
		})
	})

	Describe("Singleton", func() {
		It("should yield a single element", func() {
			s := iter.Singleton(69)

			Expect(s).To(ConsistOf(69))
		})
	})

	Describe("Skip", func() {
		It("should skip 0 elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Skip(s, 0)

			Expect(r).To(ConsistOf(69, 420))
		})

		It("should skip the given number of elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Skip(s, 1)

			Expect(r).To(ConsistOf(420))
		})

		It("should skip multiple elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Skip(s, 2)

			Expect(r).To(BeEmpty())
		})
	})

	Describe("Take", func() {
		It("should take 0 elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Take(s, 0)

			Expect(r).To(BeEmpty())
		})

		It("should take the given number of elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Take(s, 1)

			Expect(r).To(ConsistOf(69))
		})

		It("should take multiple elements", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Take(s, 2)

			Expect(r).To(ConsistOf(69, 420))
		})
	})

	Describe("Values", func() {
		It("should create a new sequence", func() {
			s := iter.Values(1, 2, 3, 4)

			Expect(s).To(ConsistOf(1, 2, 3, 4))
		})
	})
})
