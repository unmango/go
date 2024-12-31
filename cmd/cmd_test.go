package cmd_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Cmd", func() {
	It("should fail how I want it to", func() {
		cmd := exec.Command(cmdPath, "The message")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(1))
		Expect(ses.Err).To(gbytes.Say("The message\n"))
		gexec.CleanupBuildArtifacts()
	})
})
