package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("init", func() {
	Describe("version", func() {
		It("should work", Pending, func() {
			cmd := exec.Command(cmdPath, "init", "version", "v0.0.69", "--name", "blah")
			cmd.Dir = GinkgoT().TempDir()

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
		})
	})
})
