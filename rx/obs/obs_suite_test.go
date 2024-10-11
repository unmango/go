package obsv_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestObs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Obs Suite")
}
