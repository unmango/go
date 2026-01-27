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

	Describe("Map", func() {
		It("should map Some value", func() {
			m := maybe.Some(42)

			mapped := maybe.Map(m, func(v int) int {
				return v * 2
			})

			v := mapped()
			Expect(*v).To(Equal(84))
		})

		It("should not map None value", func() {
			m := maybe.None[int]()

			mapped := maybe.Map(m, func(v int) int {
				return v * 2
			})

			v := mapped()
			Expect(v).To(BeNil())
		})
	})

	Describe("Bind", func() {
		It("should bind Some value", func() {
			m := maybe.Some(42)

			bound := maybe.Bind(m, func(v int) maybe.Maybe[int] {
				return maybe.Some(v * 2)
			})

			v := bound()
			Expect(*v).To(Equal(84))
		})

		It("should not bind None value", func() {
			m := maybe.None[int]()

			bound := maybe.Bind(m, func(v int) maybe.Maybe[int] {
				return maybe.Some(v * 2)
			})

			v := bound()
			Expect(v).To(BeNil())
		})

		It("should propagate None from bind function", func() {
			m := maybe.Some(42)

			bound := maybe.Bind(m, func(v int) maybe.Maybe[int] {
				return maybe.None[int]()
			})

			v := bound()
			Expect(v).To(BeNil())
		})
	})
})
