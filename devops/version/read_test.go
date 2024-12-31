package version_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/go/devops/version"
)

var _ = Describe("Read", func() {
	Describe("ReadFile", func() {
		It("should read happypath", func() {
			v, err := version.ReadFile("test",
				version.WithRoot("testdata/happypath"),
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("0.0.69"))
		})

		It("should error when dependency does not exist", func() {
			_, err := version.ReadFile("wat",
				version.WithRoot("testdata/happypath"),
			)

			Expect(err).To(MatchError("dependency not found: wat"))
		})

		It("should error when root does not exist", func() {
			_, err := version.ReadFile("test",
				version.WithRoot("testdata/does-not-exist"),
			)

			Expect(err).To(MatchError("dependency not found: test"))
		})
	})
})
