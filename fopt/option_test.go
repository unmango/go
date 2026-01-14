package fopt_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fopt"
)

type mopts struct {
	textField   string
	numberField int
}

type (
	M func(*mopts)
	I func(mopts) mopts
)

var _ = Describe("Option", func() {
	Context("Mutable", func() {
		It("should apply", func() {
			options := &mopts{}

			fopt.Apply(options, func(o *mopts) {
				o.textField = "expected"
			})

			Expect(options.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := &mopts{}

			fopt.Apply(options,
				func(o *mopts) {
					o.textField = "expected"
				},
				func(o *mopts) {
					o.numberField = 69
				},
			)

			Expect(options.textField).To(Equal("expected"))
			Expect(options.numberField).To(Equal(69))
		})
	})

	Context("Immutable", func() {
		It("should apply", func() {
			options := mopts{}

			actual := fopt.With(options, func(o mopts) mopts {
				o.textField = "expected"
				return o
			})

			Expect(options.textField).To(BeEmpty())
			Expect(actual.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := mopts{}

			actual := fopt.With(options,
				func(o mopts) mopts {
					o.textField = "expected"
					return o
				},
				func(o mopts) mopts {
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
			op = fopt.Mut[mopts, *mopts, I](func(o *mopts) {
				o.textField = "expected"
			})
		})

		It("should apply", func() {
			options := mopts{}

			actual := fopt.With(options, op)

			Expect(options.textField).To(BeEmpty())
			Expect(actual.textField).To(Equal("expected"))
		})

		It("should apply all", func() {
			options := mopts{}

			actual := fopt.With(options,
				op,
				func(o mopts) mopts {
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
