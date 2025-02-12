package maps_test

import (
	gmaps "maps"
	gslices "slices"
	"testing/quick"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/maps"
)

var _ = Describe("Adapters", func() {
	Describe("All", func() {
		It("should match the result of maps", func() {
			fn := func(m map[string]string) bool {
				a := gmaps.All(m)
				b := maps.All(m)

				as := gmaps.Collect(a)
				bs := gmaps.Collect(b)

				return gmaps.Equal(as, bs)
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})
	})

	Describe("Collect", func() {
		It("should create a map", func() {
			var seq iter.Seq2[string, int] = func(yield func(string, int) bool) {
				yield("test", 69)
			}

			Expect(maps.Collect(seq)).To(HaveKeyWithValue("test", 69))
		})

		It("should match the result of maps", func() {
			fn := func(m map[string]string) bool {
				seq := gmaps.All(m)
				a := gmaps.Collect(seq)
				b := maps.Collect(seq)

				return gmaps.Equal(a, b)
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})
	})

	Describe("Keys", func() {
		It("should match the result of maps", func() {
			fn := func(m map[string]string) bool {
				a := gmaps.Keys(m)
				b := maps.Keys(m)

				as := gslices.Collect(a)
				bs := gslices.Collect(b)

				less := func(a, b string) bool { return a < b }
				return cmp.Equal(as, bs, cmpopts.SortSlices(less))
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})
	})

	Describe("Values", func() {
		It("should match the result of maps", func() {
			fn := func(m map[string]string) bool {
				a := gmaps.Values(m)
				b := maps.Values(m)

				as := gslices.Collect(a)
				bs := gslices.Collect(b)

				less := func(a, b string) bool { return a < b }
				return cmp.Equal(as, bs, cmpopts.SortSlices(less))
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})
	})
})
