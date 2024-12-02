package maps_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/maps"
	"github.com/unmango/go/slices"
)

var _ = Describe("Maps", func() {
	It("should append to a sequence", func() {
		seq := iter.Singleton2(69, "420")
		values := map[int]string{420: "69"}

		r := maps.AppendSeq(seq, values)

		a, b := slices.Collect2(r)
		Expect(a).To(ConsistOf(69, 420))
		Expect(b).To(ConsistOf("420", "69"))
	})
})
