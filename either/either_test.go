package either_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/either"
)

// Comprehensive tests for Either[L,R] type with pointer semantics
var _ = Describe("Either", func() {
	Describe("Left", func() {
		It("should return the given value", func() {
			e := either.Left[int](69)

			a, b := e()

			Expect(*a).To(Equal(69))
			Expect(b).To(BeNil())
		})
	})

	Describe("Right", func() {
		It("should return the given value", func() {
			e := either.Right[int](420)

			a, b := e()

			Expect(a).To(BeNil())
			Expect(*b).To(Equal(420))
		})
	})

	Describe("MapLeft", func() {
		It("should map the left value", func() {
			e := either.Left[int](69)

			r := either.MapLeft(e, func(a int) int {
				return a + 420
			})

			a, b := r()
			Expect(*a).To(Equal(489))
			Expect(b).To(BeNil())
		})

		It("should not map when left is nil", func() {
			e := either.Right[int](420)

			r := either.MapLeft(e, func(a int) int {
				return a + 1000
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(Equal(420))
		})
	})

	Describe("MapRight", func() {
		It("should map the right value", func() {
			e := either.Right[int](420)

			r := either.MapRight(e, func(b int) int {
				return b + 69
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(Equal(489))
		})

		It("should not map when right is nil", func() {
			e := either.Left[int](69)

			r := either.MapRight(e, func(b int) int {
				return b + 1000
			})

			a, b := r()
			Expect(*a).To(Equal(69))
			Expect(b).To(BeNil())
		})
	})

	Describe("BindLeft", func() {
		It("should bind the left value", func() {
			e := either.Left[int](69)

			r := either.BindLeft(e, func(a int) either.Either[int, int] {
				return either.Left[int](a + 420)
			})

			a, b := r()
			Expect(*a).To(Equal(489))
			Expect(b).To(BeNil())
		})

		It("should return right when left is nil", func() {
			e := either.Right[int](420)

			r := either.BindLeft(e, func(a int) either.Either[int, int] {
				return either.Left[int](a + 1000)
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(Equal(420))
		})
	})

	Describe("BindRight", func() {
		It("should bind the right value", func() {
			e := either.Right[int](420)

			r := either.BindRight(e, func(b int) either.Either[int, int] {
				return either.Right[int](b + 69)
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(Equal(489))
		})

		It("should return left when right is nil", func() {
			e := either.Left[int](69)

			r := either.BindRight(e, func(b int) either.Either[int, int] {
				return either.Right[int](b + 1000)
			})

			a, b := r()
			Expect(*a).To(Equal(69))
			Expect(b).To(BeNil())
		})
	})
})
