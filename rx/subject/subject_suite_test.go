package subject_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSubject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subject Suite")
}
