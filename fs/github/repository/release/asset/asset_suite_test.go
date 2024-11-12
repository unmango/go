package asset_test

import (
	"os"
	"testing"

	"github.com/google/go-github/v66/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var client *github.Client

func TestRelease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Asset Suite")
}

var _ = BeforeSuite(func() {
	client = github.NewClient(nil)

	if token, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		client = client.WithAuthToken(token)
	}
})
