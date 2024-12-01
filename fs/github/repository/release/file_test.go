package release_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/go/fs/github/repository/release"
)

var _ = Describe("File", func() {
	It("should be readonly", func() {
		fs := release.NewFs(client, "UnstoppableMango", "tdl")
		file, err := fs.Open("v0.0.29")
		Expect(err).NotTo(HaveOccurred())

		_, err = file.Write([]byte{})
		Expect(err).To(MatchError("read-only file system"))
		_, err = file.WriteAt([]byte{}, 69)
		Expect(err).To(MatchError("read-only file system"))
		_, err = file.WriteString("doesn't matter")
		Expect(err).To(MatchError("read-only file system"))
	})
})
