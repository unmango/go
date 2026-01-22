package iter_test

import (
	"maps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

type bindStub struct {
	seq iter.Seq[int]
}

var _ = Describe("Seq", func() {
	Describe("Append", func() {
		It("should append a value", func() {
			seq := slices.Values([]string{"a"})

			seq = iter.Append(seq, "b")

			Expect(seq).To(HaveExactElements("a", "b"))
		})

		It("should ignore nil seqs", func() {
			seq := iter.Append(nil, "a")

			Expect(seq).To(HaveExactElements("a"))
		})

		It("should append all values", func() {
			seq := iter.Append(nil, "a", "b")

			Expect(seq).To(HaveExactElements("a", "b"))
		})
	})

	Describe("Bind", func() {
		It("should bind", func() {
			a := iter.Singleton(69)
			b := iter.Singleton(bindStub{a})

			r := iter.Bind(b, func(s bindStub) iter.Seq[int] {
				return s.seq
			})

			Expect(r).To(ConsistOf(69))
		})

		It("should handle early termination from inner sequence", func() {
			a := slices.Values([]int{1, 2, 3})
			b := iter.Singleton(bindStub{a})

			r := iter.Bind(b, func(s bindStub) iter.Seq[int] {
				return s.seq
			})

			result := iter.Take(r, 2)
			Expect(result).To(HaveExactElements(1, 2))
		})
	})

	Describe("Compact", func() {
		It("should compact", func() {
			seq := slices.Values([]int{1, 2, 2, 3})

			res := iter.Compact(seq)

			Expect(res).To(HaveExactElements(1, 2, 3))
		})
	})

	Describe("CompactFunc", func() {
		It("should compact", func() {
			seq := slices.Values([]int{1, 2, 2, 3})

			res := iter.CompactFunc(seq, func(a, b int) bool {
				return a == b
			})

			Expect(res).To(HaveExactElements(1, 2, 3))
		})
	})

	Describe("Concat", func() {
		It("should concatenate sequences", func() {
			a := slices.Values([]int{1, 2})
			b := slices.Values([]int{3, 4})

			r := iter.Concat(a, b)

			Expect(r).To(HaveExactElements(1, 2, 3, 4))
		})

		It("should handle early termination in first sequence", func() {
			a := slices.Values([]int{1, 2, 3})
			b := slices.Values([]int{4, 5})

			r := iter.Concat(a, b)
			result := iter.Take(r, 2)

			Expect(result).To(HaveExactElements(1, 2))
		})

		It("should handle early termination in second sequence", func() {
			a := slices.Values([]int{1, 2})
			b := slices.Values([]int{3, 4, 5})

			r := iter.Concat(a, b)
			result := iter.Take(r, 3)

			Expect(result).To(HaveExactElements(1, 2, 3))
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

		It("should handle early termination", func() {
			s := slices.Values([]int{1, 2, 3, 4, 5})

			r := iter.Filter(s, func(i int) bool {
				return i > 1
			})
			result := iter.Take(r, 2)

			Expect(result).To(HaveExactElements(2, 3))
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

	Describe("All", func() {
		It("should create a sequence with the index", func() {
			seq := slices.Values([]int{2})

			res := iter.All(seq)

			Expect(maps.Collect(res)).To(Equal(map[int]int{
				0: 2,
			}))
		})

		It("should handle early termination", func() {
			seq := slices.Values([]int{1, 2, 3, 4})

			res := iter.All(seq)
			taken := iter.Take2(res, 2)

			a, b := slices.Collect2(taken)
			Expect(a).To(HaveExactElements(0, 1))
			Expect(b).To(HaveExactElements(1, 2))
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

		It("should handle early termination", func() {
			a := slices.Values([]int{1, 2, 3})
			b := slices.Values([]int{4, 5, 6})
			c := slices.Values([]iter.Seq[int]{a, b})

			r := iter.Flat(c)
			result := iter.Take(r, 4)

			Expect(result).To(HaveExactElements(1, 2, 3, 4))
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

		It("should handle early termination", func() {
			a := slices.Values([]int{1, 2, 3})
			b := slices.Values([]int{4, 5, 6})
			c := slices.Values([]iter.Seq[int]{a, b})

			r := iter.FlatMap(c, func(i int) int {
				return i * 2
			})
			result := iter.Take(r, 4)

			Expect(result).To(HaveExactElements(2, 4, 6, 8))
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

	Describe("Map", func() {
		It("should map", func() {
			s := slices.Values([]int{69, 420})

			r := iter.Map(s, func(x int) int {
				return x + 1
			})

			Expect(r).To(ConsistOf(70, 421))
		})

		It("should handle early termination", func() {
			s := slices.Values([]int{1, 2, 3, 4})

			r := iter.Map(s, func(x int) int {
				return x * 2
			})
			result := iter.Take(r, 2)

			Expect(result).To(HaveExactElements(2, 4))
		})
	})

	Describe("Pull", func() {
		It("should pull", func() {
			seq := iter.Singleton("a")

			next, _ := iter.Pull(seq)
			elem, ok := next()

			Expect(ok).To(BeTrue())
			Expect(elem).To(Equal("a"))
		})
	})

	Describe("Remove", func() {
		It("should remove the given element", func() {
			s := slices.Values([]int{1, 2, 3, 4})

			r := iter.Remove(s, 3)

			Expect(r).To(ConsistOf(1, 2, 4))
		})

		It("should handle early termination", func() {
			s := slices.Values([]int{1, 2, 3, 4, 5})

			r := iter.Remove(s, 3)
			result := iter.Take(r, 2)

			Expect(result).To(HaveExactElements(1, 2))
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

		It("should handle early termination", func() {
			s := slices.Values([]int{1, 2, 3, 4, 5})

			r := iter.Skip(s, 2)
			result := iter.Take(r, 2)

			Expect(result).To(HaveExactElements(3, 4))
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

		It("should handle early termination before limit", func() {
			s := slices.Values([]int{1, 2, 3, 4, 5})

			r := iter.Take(s, 5)
			result := iter.Take(r, 3)

			Expect(result).To(HaveExactElements(1, 2, 3))
		})
	})

	Describe("Values", func() {
		It("should create a new sequence", func() {
			s := iter.Values(1, 2, 3, 4)

			Expect(s).To(ConsistOf(1, 2, 3, 4))
		})

		It("should handle early termination", func() {
			s := iter.Values(1, 2, 3, 4, 5)

			result := iter.Take(s, 3)

			Expect(result).To(HaveExactElements(1, 2, 3))
		})
	})
})
