package either_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/either"
)

var _ = Describe("Either", func() {
	Describe("From", func() {
		It("should return the given values", func() {
			e := either.From(69, 420)

			a, b := e()

			Expect(a).To(Equal(69))
			Expect(b).To(Equal(420))
		})
	})

	Describe("Left", func() {
		It("should return the given value", func() {
			e := either.Left[int, int](69)

			a, b := e()

			Expect(a).To(Equal(69))
			Expect(b).To(Equal(0))
		})
	})

	Describe("Right", func() {
		It("should return the given value", func() {
			e := either.Right[int](420)

			a, b := e()

			Expect(a).To(Equal(0))
			Expect(b).To(Equal(420))
		})
	})

	Describe("Map", func() {
		It("should map both values", func() {
			e := either.From(69, 420)

			r := either.Map(e, func(a, b int) (string, string) {
				return "left", "right"
			})

			a, b := r()
			Expect(a).To(Equal("left"))
			Expect(b).To(Equal("right"))
		})
	})

	Describe("MapLeft", func() {
		It("should map the left value", func() {
			e := either.Left[int, int](69)

			r := either.MapLeft(e, func(a int) int {
				return a + 420
			})

			a, b := r()
			Expect(a).To(Equal(489))
			Expect(b).To(Equal(0))
		})

		It("should not map when left is zero", func() {
			e := either.Right[int](420)

			r := either.MapLeft(e, func(a int) int {
				return a + 1000
			})

			a, b := r()
			Expect(a).To(Equal(0))
			Expect(b).To(Equal(420))
		})
	})

	Describe("MapRight", func() {
		It("should map the right value", func() {
			e := either.Right[int](420)

			r := either.MapRight(e, func(b int) int {
				return b + 69
			})

			a, b := r()
			Expect(a).To(Equal(0))
			Expect(b).To(Equal(489))
		})

		It("should not map when right is zero", func() {
			e := either.Left[int, int](69)

			r := either.MapRight(e, func(b int) int {
				return b + 1000
			})

			a, b := r()
			Expect(a).To(Equal(69))
			Expect(b).To(Equal(0))
		})
	})

	Describe("Bind", func() {
		It("should bind both values", func() {
			e := either.From(69, 420)

			r := either.Bind(e, func(a, b int) either.Either[int, int] {
				return either.From(a*2, b*2)
			})

			a, b := r()
			Expect(a).To(Equal(138))
			Expect(b).To(Equal(840))
		})
	})

	Describe("BindLeft", func() {
		It("should bind the left value", func() {
			e := either.Left[int, int](69)

			r := either.BindLeft(e, func(a int) either.Either[int, int] {
				return either.From(a+420, 0)
			})

			a, b := r()
			Expect(a).To(Equal(489))
			Expect(b).To(Equal(0))
		})

		It("should return right when left is zero", func() {
			e := either.Right[int](420)

			r := either.BindLeft(e, func(a int) either.Either[int, int] {
				return either.From(a+1000, 0)
			})

			a, b := r()
			Expect(a).To(Equal(0))
			Expect(b).To(Equal(420))
		})
	})

	Describe("BindRight", func() {
		It("should bind the right value", func() {
			e := either.Right[int](420)

			r := either.BindRight(e, func(b int) either.Either[int, int] {
				return either.From(0, b+69)
			})

			a, b := r()
			Expect(a).To(Equal(0))
			Expect(b).To(Equal(489))
		})

		It("should return left when right is zero", func() {
			e := either.Left[int, int](69)

			r := either.BindRight(e, func(b int) either.Either[int, int] {
				return either.From(0, b+1000)
			})

			a, b := r()
			Expect(a).To(Equal(69))
			Expect(b).To(Equal(0))
		})
	})
})
