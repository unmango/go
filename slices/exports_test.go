package slices_test

import (
	"testing/quick"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	goslices "slices"

	"github.com/unmango/go/slices"
)

var _ = Describe("Collect", func() {
	It("should equal slices.Collect", func() {
		f := func(s []int) bool {
			seq := func(yield func(int) bool) {
				for x := range s {
					if !yield(x) {
						break
					}
				}
			}

			a := slices.Collect(seq)
			b := goslices.Collect(seq)

			return goslices.Equal(a, b)
		}

		Expect(quick.Check(f, nil)).To(Succeed())
	})

	It("should equal slices.Compact", func() {
		f := func(s []int) bool {
			slice := []int{69}

			a := slices.Compact(slice)
			b := goslices.Compact(slice)

			return goslices.Equal(a, b)
		}

		Expect(quick.Check(f, nil)).To(Succeed())
	})

	It("should equal slices.CompactFunc", func() {
		f := func(s []int) bool {
			slice := []int{69}
			eq := func(a, b int) bool {
				return a == b
			}

			a := slices.CompactFunc(slice, eq)
			b := goslices.CompactFunc(slice, eq)

			return goslices.Equal(a, b)
		}

		Expect(quick.Check(f, nil)).To(Succeed())
	})
})
