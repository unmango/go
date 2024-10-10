package maps_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/maps"
)

var _ = Describe("Collect", func() {
	It("should create a map", func() {
		var seq iter.Seq2[string, int] = func(yield func(string, int) bool) {
			yield("test", 69)
		}

		Expect(maps.Collect(seq)).To(HaveKeyWithValue("test", 69))
	})
})
