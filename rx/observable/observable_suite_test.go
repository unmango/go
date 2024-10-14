package obs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestObservable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observable Suite")
}
