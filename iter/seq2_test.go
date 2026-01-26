package iter_test

import (
	"maps"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Seq2", func() {
	Describe("Append", func() {
		It("should append", func() {
			seq := iter.Singleton2("a", 1)

			seq = iter.Append2(seq, "b", 2)

			Expect(maps.Collect(seq)).To(Equal(map[string]int{
				"a": 1,
				"b": 2,
			}))
		})

		It("should ignore nil seqs", func() {
			seq := iter.Append2(nil, "a", 1)

			Expect(maps.Collect(seq)).To(Equal(map[string]int{
				"a": 1,
			}))
		})

		It("should handle early termination", func() {
			seq := iter.Singleton2("a", 1)
			seq = iter.Append2(seq, "b", 2)
			seq = iter.Append2(seq, "c", 3)

			result := iter.Take2(seq, 2)

			a, b := slices.Collect2(result)
			Expect(len(a)).To(Equal(2))
			Expect(len(b)).To(Equal(2))
		})

		It("should handle early termination before appended element", func() {
			base := slices.Zip([]string{"a", "b", "c"}, []int{1, 2, 3})
			seq := iter.Append2(base, "d", 4)

			result := iter.Take2(seq, 2)

			a, b := slices.Collect2(result)
			Expect(a).To(HaveExactElements("a", "b"))
			Expect(b).To(HaveExactElements(1, 2))
		})
	})

	Describe("Bind2", func() {
		It("should bind sequences", func() {
			seq := slices.Zip([]int{1, 2}, []string{"a", "b"})

			result := iter.Bind2(seq, func(k int, v string) iter.Seq2[int, string] {
				return iter.Singleton2(k*2, v+v)
			})

			a, b := slices.Collect2(result)
			Expect(a).To(ConsistOf(2, 4))
			Expect(b).To(ConsistOf("aa", "bb"))
		})

		It("should handle early termination", func() {
			seq := slices.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})

			result := iter.Bind2(seq, func(k int, v string) iter.Seq2[int, string] {
				return slices.Zip([]int{k, k * 2}, []string{v, v + v})
			})

			taken := iter.Take2(result, 3)

			a, b := slices.Collect2(taken)
			Expect(len(a)).To(Equal(3))
			Expect(len(b)).To(Equal(3))
		})

		It("should handle empty inner sequence", func() {
			seq := slices.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})

			result := iter.Bind2(seq, func(k int, v string) iter.Seq2[int, string] {
				if k == 2 {
					return iter.Empty2[int, string]()
				}
				return iter.Singleton2(k*2, v+v)
			})

			a, b := slices.Collect2(result)
			Expect(a).To(ConsistOf(2, 6))
			Expect(b).To(ConsistOf("aa", "cc"))
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

	Describe("Empty", func() {
		It("should not yield any elements", func() {
			seq := iter.Empty2[int, string]()

			a, b := slices.Collect2(seq)
			Expect(a).To(BeEmpty())
			Expect(b).To(BeEmpty())
		})
	})

	Describe("Fold", func() {
		It("should fold", func() {
			s := slices.Zip([]int{69, 420}, []string{"69", "420"})

			r := iter.Fold2(s, func(sum int, i int, _ string) int {
				return sum + i
			}, 0)

			Expect(r).To(Equal(489))
		})
	})

	Describe("Filter", func() {
		It("should not yield the filtered element", func() {
			s := slices.Zip([]int{69, 420}, []string{"69", "420"})

			r := iter.Filter2(s, func(i int, _ string) bool {
				return i == 69
			})

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("69"))
		})

		It("should handle early termination", func() {
			s := slices.Zip([]int{1, 2, 3, 4}, []string{"a", "b", "c", "d"})

			r := iter.Filter2(s, func(i int, _ string) bool {
				return i > 1
			})
			result := iter.Take2(r, 2)

			a, b := slices.Collect2(result)
			Expect(a).To(HaveExactElements(2, 3))
			Expect(b).To(HaveExactElements("b", "c"))
		})
	})

	Describe("Head2", func() {
		It("should return the first element", func() {
			seq := slices.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})

			k, v := iter.Head2(seq)

			Expect(k).To(Equal(1))
			Expect(v).To(Equal("a"))
		})

		It("should return zero values for empty sequence", func() {
			seq := iter.Empty2[int, string]()

			k, v := iter.Head2(seq)

			Expect(k).To(Equal(0))
			Expect(v).To(Equal(""))
		})
	})

	Describe("Skip2", func() {
		It("should skip elements", func() {
			seq := slices.Zip([]int{1, 2, 3, 4}, []string{"a", "b", "c", "d"})

			result := iter.Skip2(seq, 2)

			a, b := slices.Collect2(result)
			Expect(a).To(HaveExactElements(3, 4))
			Expect(b).To(HaveExactElements("c", "d"))
		})

		It("should handle early termination", func() {
			seq := slices.Zip([]int{1, 2, 3, 4, 5}, []string{"a", "b", "c", "d", "e"})

			result := iter.Skip2(seq, 2)
			taken := iter.Take2(result, 2)

			a, b := slices.Collect2(taken)
			Expect(a).To(HaveExactElements(3, 4))
			Expect(b).To(HaveExactElements("c", "d"))
		})
	})

	Describe("Take2", func() {
		It("should take elements", func() {
			seq := slices.Zip([]int{1, 2, 3, 4}, []string{"a", "b", "c", "d"})

			result := iter.Take2(seq, 2)

			a, b := slices.Collect2(result)
			Expect(a).To(HaveExactElements(1, 2))
			Expect(b).To(HaveExactElements("a", "b"))
		})

		It("should handle early termination before limit", func() {
			seq := slices.Zip([]int{1, 2, 3, 4, 5}, []string{"a", "b", "c", "d", "e"})

			result := iter.Take2(seq, 10)
			taken := iter.Take2(result, 3)

			a, b := slices.Collect2(taken)
			Expect(a).To(HaveExactElements(1, 2, 3))
			Expect(b).To(HaveExactElements("a", "b", "c"))
		})
	})

	Describe("Map", func() {
		It("should map", func() {
			s := slices.Zip([]int{69, 420}, []string{"69", "420"})

			r := iter.Map2(s, func(i int, s string) (int, string) {
				return i + 1, s
			})

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf(70, 421))
			Expect(b).To(ConsistOf("69", "420"))
		})
	})

	Describe("Pull", func() {
		It("should pull", func() {
			seq := iter.Singleton2("a", 1)

			next, _ := iter.Pull2(seq)
			a, b, ok := next()

			Expect(ok).To(BeTrue())
			Expect(a).To(Equal("a"))
			Expect(b).To(Equal(1))
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
})
