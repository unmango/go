package release_test

import (
	"os"
	"testing"

	"github.com/google/go-github/v67/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var client *github.Client

func TestRelease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Release Suite")
}

var _ = BeforeSuite(func() {
	client = github.NewClient(nil)

	if token, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		client = client.WithAuthToken(token)
	}
})
