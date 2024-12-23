package user_test

import (
	"testing"

	"github.com/google/go-github/v68/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/go/fs/github/internal"
)

var client *github.Client

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = BeforeSuite(func() {
	client = internal.DefaultClient()
})
