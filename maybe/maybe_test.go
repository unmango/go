package maybe_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/maybe"
)

var _ = Describe("Maybe", func() {
	Describe("None", func() {
		It("should be None", func() {
			Expect(maybe.None).To(Satisfy(maybe.IsNone[any]))
		})

		It("should not be ok", func() {
			Expect(maybe.None).NotTo(Satisfy(maybe.IsOk[any]))
		})

		It("should be nil", func() {
			Expect(maybe.None()).To(BeNil())
		})

		It("should be ErrNone", func() {
			_, err := maybe.None()

			Expect(err).To(BeIdenticalTo(maybe.ErrNone))
		})
	})

	Describe("Ok", func() {
		DescribeTable("IsNone",
			Entry("number should not be none", 1),
			Entry("string should not be none", "thing"),
			Entry("char should not be none", 't'),
			Entry("false should not be none", false),
			Entry("true should not be none", true),
			func(x any) {
				Expect(maybe.Ok(x)).NotTo(Satisfy(maybe.IsNone[any]))
			},
		)
	})
})
