package testing

import (
	"context"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func GitInit(ctx context.Context, path string) {
	GinkgoHelper()

	cmd := exec.CommandContext(ctx, "git", "init", path)
	ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

	Expect(err).NotTo(HaveOccurred())
	Eventually(ses).Should(gexec.Exit(0))
}

func TempGitRepo(t T) string {
	GinkgoHelper()

	path := t.TempDir()
	GitInit(context.Background(), path)
	return path
}
