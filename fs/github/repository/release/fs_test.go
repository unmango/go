package release_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github/repository/release"
)

var _ = Describe("Fs", func() {
	It("should stat", func() {
		fs := release.NewFs(client, "UnstoppableMango", "tdl")

		r, err := fs.Stat("v0.0.29")

		Expect(err).NotTo(HaveOccurred())
		Expect(r.Name()).To(Equal("v0.0.29"))
	})
})
