package either_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/either"
)

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
		It("should map the given value", func() {
			e := either.Left[int](69)

			r := either.MapLeft(e, func(a int) int {
				return a + 420
			})

			a, b := r()
			Expect(*a).To(Equal(489))
			Expect(b).To(BeNil())
		})
	})

	Describe("MapRight", func() {
		It("should map the given value", func() {
			e := either.Right[int](420)

			r := either.MapRight(e, func(b int) int {
				return b + 69
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(Equal(489))
		})
	})

	Describe("BindLeft", func() {
		It("should bind the given value", func() {
			e := either.Left[int](69)

			r := either.BindLeft(e, func(a int) either.Either[int, int] {
				return either.Left[int](a + 420)
			})

			a, b := r()
			Expect(*a).To(Equal(489))
			Expect(b).To(BeNil())
		})
	})

	Describe("BindRight", func() {
		It("should bind the given value", func() {
			e := either.Right[int](420)

			r := either.BindRight(e, func(b int) either.Either[int, int] {
				return either.Right[int](b + 69)
			})

			a, b := r()
			Expect(a).To(BeNil())
			Expect(*b).To(BeEquivalentTo(489))
		})
	})
})
