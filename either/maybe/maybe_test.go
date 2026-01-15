package maybe_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/either/maybe"
)

var _ = Describe("Maybe", func() {
	Describe("Some", func() {
		It("should be Some", func() {
			m := maybe.Some(69)

			v := m()

			Expect(*v).To(Equal(69))
		})
	})

	Describe("None", func() {
		It("should be None", func() {
			m := maybe.None[any]()

			v := m()

			Expect(v).To(BeNil())
		})
	})
})
