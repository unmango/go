package user_test

import (
	"github.com/google/go-github/v66/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github/user"
)

var _ = Describe("File", func() {
	It("should list repositories", func() {
		client := github.NewClient(nil)
		fs := user.NewFs(client)
		file, err := fs.Open("UnstoppableMango")
		Expect(err).NotTo(HaveOccurred())

		repos, err := file.Readdir(69)

		Expect(err).NotTo(HaveOccurred())
		Expect(repos).NotTo(BeEmpty())
		names := make([]string, len(repos))
		for i, r := range repos {
			names[i] = r.Name()
		}
		Expect(names).To(ContainElement("advent-of-code"))
	})
})
