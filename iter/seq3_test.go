package iter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
)

var _ = Describe("Seq3", func() {
	Describe("Bind3", func() {
		It("should bind sequences", func() {
			seq := slices.Values3([]int{1, 2}, []string{"a", "b"}, []bool{true, false})

			result := iter.Bind3(seq, func(i int, s string, b bool) iter.Seq3[int, string, bool] {
				return iter.Singleton3(i*2, s+s, !b)
			})

			a, b, c := slices.Collect3(result)
			Expect(a).To(ConsistOf(2, 4))
			Expect(b).To(ConsistOf("aa", "bb"))
			Expect(c).To(ConsistOf(false, true))
		})

		It("should handle early termination in inner sequence", func() {
			seq := iter.Singleton3(1, "a", true)

			result := iter.Bind3(seq, func(i int, s string, b bool) iter.Seq3[int, string, bool] {
				return slices.Values3(
					[]int{i, i * 2, i * 3},
					[]string{s, s + s, s + s + s},
					[]bool{b, !b, b},
				)
			})

			a, b, c := slices.Collect3(result)
			Expect(a).To(HaveExactElements(1, 2, 3))
			Expect(b).To(HaveExactElements("a", "aa", "aaa"))
			Expect(c).To(HaveExactElements(true, false, true))
		})

		It("should handle empty inner sequence", func() {
			seq := slices.Values3([]int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

			result := iter.Bind3(seq, func(i int, s string, b bool) iter.Seq3[int, string, bool] {
				if i == 2 {
					return iter.Empty3[int, string, bool]()
				}
				return iter.Singleton3(i*2, s+s, !b)
			})

			a, b, c := slices.Collect3(result)
			Expect(a).To(ConsistOf(2, 6))
			Expect(b).To(ConsistOf("aa", "cc"))
			Expect(c).To(ConsistOf(false, false))
		})
	})

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

	Describe("Filter", func() {
		It("should not yield the filtered element", func() {
			seq := slices.Values3(
				[]int{69, 420},
				[]string{"69", "420"},
				[]bool{true, true},
			)

			r := iter.Filter3(seq, func(i int, _ string, _ bool) bool {
				return i == 69
			})

			a, b, c := slices.Collect3(r)
			Expect(a).To(ConsistOf(69))
			Expect(b).To(ConsistOf("69"))
			Expect(c).To(ConsistOf(true))
		})
	})

	Describe("Fold", func() {
		It("should fold", func() {
			seq := slices.Values3(
				[]int{69, 420},
				[]string{"69", "420"},
				[]bool{true, true},
			)

			r := iter.Fold3(seq, func(acc int, i int, _ string, _ bool) int {
				return acc + i
			}, 0)

			Expect(r).To(Equal(489))
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

	Describe("Map", func() {
		It("should map", func() {
			seq := slices.Values3(
				[]int{69, 420},
				[]string{"69", "420"},
				[]bool{true, true},
			)

			r := iter.Map3(seq, func(i int, s string, b bool) (int, string, bool) {
				return i + 1, s, b
			})

			a, b, c := slices.Collect3(r)
			Expect(a).To(ConsistOf(70, 421))
			Expect(b).To(ConsistOf("69", "420"))
			Expect(c).To(ConsistOf(true, true))
		})
	})
})
