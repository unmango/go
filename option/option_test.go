package option_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/option"
)

type MockOptions struct {
	textField   string
	numberField int
}

type (
	M func(*MockOptions)
	I func(MockOptions) MockOptions
)

var _ = Describe("Option", func() {
	Context("Mutable", func() {
		It("should apply", func() {
			options := &MockOptions{}

			option.Apply(options, func(o *MockOptions) {
				o.textField = "expected"
			})

			Expect(options.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := &MockOptions{}

			option.Apply(options,
				func(o *MockOptions) {
					o.textField = "expected"
				},
				func(o *MockOptions) {
					o.numberField = 69
				},
			)

			Expect(options.textField).To(Equal("expected"))
			Expect(options.numberField).To(Equal(69))
		})
	})

	Context("Immutable", func() {
		It("should apply", func() {
			options := MockOptions{}

			actual := option.With(options, func(o MockOptions) MockOptions {
				o.textField = "expected"
				return o
			})

			Expect(options.textField).To(BeEmpty())
			Expect(actual.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := MockOptions{}

			actual := option.With(options,
				func(o MockOptions) MockOptions {
					o.textField = "expected"
					return o
				},
				func(o MockOptions) MockOptions {
					o.numberField = 69
					return o
				},
			)

			Expect(options.textField).To(BeEmpty())
			Expect(options.numberField).To(BeZero())
			Expect(actual.textField).To(Equal("expected"))
			Expect(actual.numberField).To(Equal(69))
		})
	})

	Context("Mut", func() {
		var op I

		BeforeEach(func() {
			op = option.Mut[MockOptions, *MockOptions, I](func(o *MockOptions) {
				o.textField = "expected"
			})
		})

		It("should apply", func() {
			options := MockOptions{}

			actual := option.With(options, op)

			Expect(options.textField).To(BeEmpty())
			Expect(actual.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := MockOptions{}

			actual := option.With(options,
				op,
				func(o MockOptions) MockOptions {
					o.numberField = 69
					return o
				},
			)

			Expect(options.textField).To(BeEmpty())
			Expect(options.numberField).To(BeZero())
			Expect(actual.textField).To(Equal("expected"))
			Expect(actual.numberField).To(Equal(69))
		})
	})
})
