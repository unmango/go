package user_test

import (
	"os"

	"github.com/google/go-github/v67/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github/user"
)

var _ = Describe("Fs", func() {
	It("should open user", func() {
		client := github.NewClient(nil)
		fs := user.NewFs(client)

		user, err := fs.Open("UnstoppableMango")

		Expect(err).NotTo(HaveOccurred())
		Expect(user.Name()).To(Equal("UnstoppableMango"))
	})

	It("should open user file", func() {
		client := github.NewClient(nil)
		fs := user.NewFs(client)

		user, err := fs.OpenFile("UnstoppableMango", 69, os.ModePerm)

		Expect(err).NotTo(HaveOccurred())
		Expect(user.Name()).To(Equal("UnstoppableMango"))
	})

	It("should stat user", func() {
		client := github.NewClient(nil)
		fs := user.NewFs(client)

		user, err := fs.Stat("UnstoppableMango")

		Expect(err).NotTo(HaveOccurred())
		Expect(user.Name()).To(Equal("UnstoppableMango"))
	})
})
