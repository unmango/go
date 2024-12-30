package version_test

import (
	"context"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Version", func() {
	It("should sprint", func(ctx context.Context) {
		cmd := exec.Command(cmdPath, "test")
		cmd.Dir = "./testdata/happypath"

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Consistently(ses.Err).Should(gbytes.Say(""))
		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("0.0.69"))
	})
})
