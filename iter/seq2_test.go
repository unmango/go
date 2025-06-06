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
			s := slices.Values2([]int{69, 420}, []string{"69", "420"})

			r := iter.Fold2(s, func(sum int, i int, _ string) int {
				return sum + i
			}, 0)

			Expect(r).To(Equal(489))
		})
	})

	Describe("Filter", func() {
		It("should not yield the filtered element", func() {
			s := slices.Values2([]int{69, 420}, []string{"69", "420"})

			r := iter.Filter2(s, func(i int, _ string) bool {
				return i == 69
			})

			a, b := slices.Collect2(r)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("69"))
		})
	})

	Describe("Map", func() {
		It("should map", func() {
			s := slices.Values2([]int{69, 420}, []string{"69", "420"})

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
