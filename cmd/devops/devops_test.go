package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("devops", func() {
	Describe("version", func() {
		It("should print the version", func() {
			cmd := exec.Command(cmdPath, "version", "test")
			cmd.Dir = "testdata/happypath"

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Out).To(gbytes.Say("0.0.69\n"))
		})

		It("should print the prefixed version", Pending, func() {
			cmd := exec.Command(cmdPath, "version", "test", "--prefixed")
			cmd.Dir = "testdata/happypath"

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Out).To(gbytes.Say("0.0.69\n"))
		})

		It("should should error when dependency does not exist", func() {
			cmd := exec.Command(cmdPath, "version", "wat")
			cmd.Dir = "testdata/happypath"

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(1))
			Expect(ses.Err).To(gbytes.Say("dependency not found: wat\n"))
		})
	})
})
