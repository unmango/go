package git_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/testing"
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
		testing.GitInit(ctx, wd)

		By("Creating a subdirectory")
		subdir := filepath.Join(wd, "subdir")
		Expect(os.Mkdir(subdir, os.ModePerm)).To(Succeed())

		p, err := git.Root(git.WithWorkingDirectory(ctx, subdir))

		Expect(err).NotTo(HaveOccurred())
		p = strings.TrimPrefix(p, "/private") // Mac crap
		Expect(p).To(Equal(wd))
	})

	It("should print the git root from the environment", func(ctx context.Context) {
		os.Setenv("GIT_ROOT", "/some/random/path")
		DeferCleanup(func() {
			os.Unsetenv("GIT_ROOT")
		})

		p, err := git.Root(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(p).To(Equal("/some/random/path"))
	})

	It("should use git from the environment", func(ctx context.Context) {
		wd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		expected := strings.TrimSuffix(wd, "/vcs/git")
		exe, err := exec.LookPath("git")
		Expect(err).NotTo(HaveOccurred())

		os.Setenv("GIT_PATH", exe)
		DeferCleanup(func() {
			os.Unsetenv("GIT_PATH")
		})

		p, err := git.Root(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(p).To(Equal(expected))
	})

	It("should use git from the context", func(ctx context.Context) {
		wd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		expected := strings.TrimSuffix(wd, "/vcs/git")
		exe, err := exec.LookPath("git")
		Expect(err).NotTo(HaveOccurred())

		p, err := git.Root(git.WithGitPath(ctx, exe))

		Expect(err).NotTo(HaveOccurred())
		Expect(p).To(Equal(expected))
	})
})
