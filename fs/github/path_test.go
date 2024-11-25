package github_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github"
)

var _ = Describe("Path", func() {
	Describe("Parse", func() {
		DescribeTable("Owner",
			Entry(nil, "unmango", "unmango"),
			Entry(nil, "https://github.com/unmango", "unmango"),
			Entry(nil, "github.com/unmango", "unmango"),
			Entry(nil, "https://api.github.com/unmango", "unmango"),
			Entry(nil, "api.github.com/unmango", "unmango"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Owner()).To(Equal(name))
			},
		)

		DescribeTable("Repository",
			Entry(nil, "unmango/go", "go"),
			Entry(nil, "https://github.com/unmango/go", "go"),
			Entry(nil, "github.com/unmango/go", "go"),
			Entry(nil, "https://api.github.com/unmango/go", "go"),
			Entry(nil, "api.github.com/unmango/go", "go"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Repository()).To(Equal(name))
			},
		)

		DescribeTable("No Repository",
			Entry(nil, "unmango"),
			Entry(nil, "https://github.com/unmango"),
			Entry(nil, "github.com/unmango"),
			Entry(nil, "https://api.github.com/unmango"),
			Entry(nil, "api.github.com/unmango"),
			func(input string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				_, err = res.Repository()
				Expect(err).To(MatchError("no repository"))
			},
		)

		DescribeTable("Branch",
			Entry(nil, "https://github.com/unmango/go/tree/main", "main"),
			Entry(nil, "https://api.github.com/unmango/go/tree/main", "main"),
			Entry(nil, "api.github.com/unmango/go/tree/main", "main"),
			Entry(nil, "github.com/unmango/go/tree/main", "main"),
			Entry(nil, "unmango/go/tree/main", "main"),
			Entry(nil, "unmango/go/tree/main/fs", "main"),
			Entry(nil, "unmango/go/tree/main/fs/path_test.go", "main"),
			Entry(nil, "unmango/go/tree/feature-name", "feature-name"),
			Entry(nil, "https://raw.githubusercontent.com/unmango/go/refs/heads/main/fs/fold.go", "main"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Branch()).To(Equal(name))
			},
		)

		DescribeTable("Not a Branch",
			Entry(nil, "unmango"),
			Entry(nil, "https://github.com/unmango"),
			Entry(nil, "github.com/unmango"),
			Entry(nil, "https://api.github.com/unmango"),
			Entry(nil, "api.github.com/unmango"),
			Entry(nil, "unmango/go"),
			Entry(nil, "https://github.com/unmango/go"),
			Entry(nil, "github.com/unmango/go"),
			Entry(nil, "https://api.github.com/unmango/go"),
			Entry(nil, "api.github.com/unmango/go"),
			func(input string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				_, err = res.Branch()
				Expect(err).To(MatchError("not a branch"))
			},
		)

		DescribeTable("No Branch",
			Entry(nil, "https://github.com/unmango/go/tree"),
			Entry(nil, "https://api.github.com/unmango/go/tree"),
			Entry(nil, "api.github.com/unmango/go/tree"),
			Entry(nil, "github.com/unmango/go/tree"),
			Entry(nil, "https://raw.githubusercontent.com/unmango/go/refs/heads"),
			Entry(nil, "https://raw.githubusercontent.com/unmango/go/refs"),
			func(input string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				_, err = res.Branch()
				Expect(err).To(MatchError("no branch"))
			},
		)

		DescribeTable("Release",
			Entry(nil, "https://github.com/unmango/go/releases/tag/v0.0.69", "v0.0.69"),
			Entry(nil, "https://api.github.com/unmango/go/releases/tag/v0.0.69", "v0.0.69"),
			Entry(nil, "api.github.com/unmango/go/releases/tag/v0.0.69", "v0.0.69"),
			Entry(nil, "github.com/unmango/go/releases/tag/v0.0.69", "v0.0.69"),
			Entry(nil, "unmango/go/releases/tag/v0.0.69", "v0.0.69"),
			Entry(nil, "https://github.com/unmango/go/releases/download/v0.0.69", "v0.0.69"),
			Entry(nil, "https://api.github.com/unmango/go/releases/download/v0.0.69", "v0.0.69"),
			Entry(nil, "api.github.com/unmango/go/releases/download/v0.0.69", "v0.0.69"),
			Entry(nil, "github.com/unmango/go/releases/download/v0.0.69", "v0.0.69"),
			Entry(nil, "unmango/go/releases/download/v0.0.69", "v0.0.69"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Release()).To(Equal(name))
			},
		)

		DescribeTable("No Release",
			Entry(nil, "https://github.com/unmango/go/releases/bleh/v0.0.69", "v0.0.69"),
			Entry(nil, "https://api.github.com/unmango/go/releases/bleh/v0.0.69", "v0.0.69"),
			Entry(nil, "api.github.com/unmango/go/releases/bleh/v0.0.69", "v0.0.69"),
			Entry(nil, "github.com/unmango/go/releases/bleh/v0.0.69", "v0.0.69"),
			Entry(nil, "unmango/go/releases/bleh/v0.0.69", "v0.0.69"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				_, err = res.Release()
				Expect(err).To(MatchError("no release"))
			},
		)

		DescribeTable("Asset",
			Entry(nil, "https://github.com/unmango/go/releases/tag/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "https://api.github.com/unmango/go/releases/tag/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "api.github.com/unmango/go/releases/tag/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "github.com/unmango/go/releases/tag/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "unmango/go/releases/tag/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "https://github.com/unmango/go/releases/download/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "https://api.github.com/unmango/go/releases/download/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "api.github.com/unmango/go/releases/download/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "github.com/unmango/go/releases/download/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			Entry(nil, "unmango/go/releases/download/v0.0.69/my-asset.tar.gz", "my-asset.tar.gz"),
			func(input, name string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Asset()).To(Equal(name))
			},
		)

		DescribeTable("No Asset",
			Entry(nil, "https://github.com/unmango/go/releases/tag/v0.0.69"),
			Entry(nil, "https://api.github.com/unmango/go/releases/tag/v0.0.69"),
			Entry(nil, "api.github.com/unmango/go/releases/tag/v0.0.69"),
			Entry(nil, "github.com/unmango/go/releases/tag/v0.0.69"),
			Entry(nil, "unmango/go/releases/tag/v0.0.69"),
			Entry(nil, "https://github.com/unmango/go/releases/download/v0.0.69"),
			Entry(nil, "https://api.github.com/unmango/go/releases/download/v0.0.69"),
			Entry(nil, "api.github.com/unmango/go/releases/download/v0.0.69"),
			Entry(nil, "github.com/unmango/go/releases/download/v0.0.69"),
			Entry(nil, "unmango/go/releases/download/v0.0.69"),
			func(input string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				_, err = res.Asset()
				Expect(err).To(MatchError("no asset"))
			},
		)

		DescribeTable("Content",
			Entry(nil, "https://github.com/unmango/go/tree/main", []string{}),
			Entry(nil, "https://api.github.com/unmango/go/tree/main", []string{}),
			Entry(nil, "api.github.com/unmango/go/tree/main", []string{}),
			Entry(nil, "github.com/unmango/go/tree/main", []string{}),
			Entry(nil, "unmango/go/tree/main", []string{}),
			Entry(nil, "unmango/go/tree/main/fs", []string{"fs"}),
			Entry(nil, "unmango/go/tree/main/fs/path_test.go", []string{"fs", "path_test.go"}),
			Entry(nil, "https://raw.githubusercontent.com/unmango/go/refs/heads/main/fs/fold.go", []string{"fs", "fold.go"}),
			func(input string, parts []string) {
				res, err := github.Parse(input)

				Expect(err).NotTo(HaveOccurred())
				Expect(res.Content()).To(Equal(parts))
			},
		)
	})
})
