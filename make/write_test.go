package make_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/make"
)

var _ = FDescribe("Write", func() {
	It("should write a line", func() {
		buf := &bytes.Buffer{}
		w := make.NewWriter(buf)

		n, err := w.WriteLine()

		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(1))
		Expect(buf.String()).To(Equal("\n"))
	})

	It("should write a target", func() {
		buf := &bytes.Buffer{}
		w := make.NewWriter(buf)

		n, err := w.WriteTarget("target")

		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(7))
		Expect(buf.String()).To(Equal("target:"))
	})

	It("should write multiple targets", func() {
		buf := &bytes.Buffer{}
		w := make.NewWriter(buf)

		n, err := w.WriteTargets([]string{"target", "target2"})

		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(15))
		Expect(buf.String()).To(Equal("target target2:"))
	})
})
