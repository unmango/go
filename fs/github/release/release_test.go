package release_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github/release"
)

var _ = Describe("Release", func() {
	It("should work", func() {
		r := release.New("UnstoppableMango", "tdl", "v0.0.29", client.Repositories)

		_, err := r.Stat("tdl-amd64-linux.tar.gz")

		Expect(err).NotTo(HaveOccurred())
	})
})
