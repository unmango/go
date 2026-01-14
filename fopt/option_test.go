package fopt_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fopt"
)

type mopts struct {
	textField   string
	numberField int
}

type (
	M  func(*mopts)
	I  func(mopts) mopts
	Ma func(*mopts) error
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

	Context("Maybe", func() {
		It("should apply without error", func() {
			options := &mopts{}

			err := fopt.TryApply(options, func(o *mopts) error {
				o.textField = "expected"
				return nil
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(options.textField).To(Equal("expected"))
		})

		It("should apply all without error", func() {
			options := &mopts{}

			err := fopt.TryApply(options,
				func(o *mopts) error {
					o.textField = "expected"
					return nil
				},
				func(o *mopts) error {
					o.numberField = 69
					return nil
				},
			)

			Expect(err).ToNot(HaveOccurred())
			Expect(options.textField).To(Equal("expected"))
			Expect(options.numberField).To(Equal(69))
		})

		It("should return error from first failing option", func() {
			options := &mopts{}
			expectedErr := errors.New("option failed")

			err := fopt.TryApply(options,
				func(o *mopts) error {
					o.textField = "expected"
					return nil
				},
				func(o *mopts) error {
					return expectedErr
				},
				func(o *mopts) error {
					o.numberField = 69
					return nil
				},
			)

			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(expectedErr))
			Expect(options.textField).To(Equal("expected"))
			Expect(options.numberField).To(BeZero())
		})

		It("should handle empty options list", func() {
			options := &mopts{}
			var emptyOptions []Ma

			err := fopt.TryApplyAll(options, emptyOptions)

			Expect(err).ToNot(HaveOccurred())
		})

		It("should return error immediately on first failure", func() {
			options := &mopts{}
			firstErr := errors.New("first option failed")

			err := fopt.TryApply(options,
				func(o *mopts) error {
					return firstErr
				},
				func(o *mopts) error {
					o.textField = "should not be set"
					return nil
				},
			)

			Expect(err).To(Equal(firstErr))
			Expect(options.textField).To(BeEmpty())
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
