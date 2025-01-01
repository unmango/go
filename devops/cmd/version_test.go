package cmd_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/devops/cmd"
	"github.com/unmango/go/vcs/git"
)

var _ = Describe("Version", func() {
	Describe("VersionOptions", func() {
		It("should return the chdir when it is defined", func(ctx context.Context) {
			o := cmd.VersionOptions{Chdir: "blah"}

			p, err := o.Cwd(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(p.Path()).To(Equal("blah"))
		})

		It("should return the git path with chdir is empty", func(ctx context.Context) {
			expected, err := git.Root(ctx)
			Expect(err).NotTo(HaveOccurred())
			o := cmd.VersionOptions{}

			p, err := o.Cwd(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(p.Path()).To(Equal(expected))
		})
	})
})
