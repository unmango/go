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
})
