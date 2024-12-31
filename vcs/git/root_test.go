package git_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"github.com/unmango/go/vcs/git"
)

var _ = Describe("Root", func() {
	It("should print the current git root", func(ctx context.Context) {
		wd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		expected := strings.TrimSuffix(wd, "/vcs/git")

		p, err := git.Root(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(p).To(Equal(expected))
	})

	It("should print the working directory's git root", func(ctx context.Context) {
		wd := GinkgoT().TempDir()
		cmd := exec.CommandContext(ctx, "git", "init")
		cmd.Dir = wd
		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))

		By("Creating a subdirectory")
		subdir := filepath.Join(wd, "subdir")
		Expect(os.Mkdir(subdir, os.ModePerm)).To(Succeed())

		p, err := git.Root(git.WithWorkingDirectory(ctx, subdir))

		Expect(err).NotTo(HaveOccurred())
		p = strings.TrimPrefix(p, "/private") // Mac crap
		Expect(p).To(Equal(wd))
	})
})
