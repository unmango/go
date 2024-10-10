package result_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/result"
)

var _ = Describe("Result", func() {
	Context("Err", func() {
		It("should return the error value", func() {
			err := errors.New("test error")

			var r result.Result[int] = func() (int, error) {
				return 69, err
			}

			Expect(r.Err()).To(BeIdenticalTo(err))
		})
	})
})
