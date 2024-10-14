package option_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/option"
)

type MockOptions struct {
	textField string
	intField  int
}

type M func(*MockOptions)

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

		})
	})
})
