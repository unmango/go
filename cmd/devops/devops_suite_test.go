package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var cmdPath string

func TestDevops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Devops Suite")
}

var _ = BeforeSuite(func() {
	p, err := gexec.Build("main.go")
	Expect(err).NotTo(HaveOccurred())
	cmdPath = p
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
