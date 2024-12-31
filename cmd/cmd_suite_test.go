package cmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var cmdPath string

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmd Suite")
}

var _ = BeforeSuite(func() {
	p, err := gexec.Build("./testdata")
	Expect(err).NotTo(HaveOccurred())
	cmdPath = p
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
