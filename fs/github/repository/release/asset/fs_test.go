package asset_test

import (
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/github/repository/release/asset"
)

var _ = Describe("Fs", func() {
	It("should stat an asset", func() {
		r := asset.NewFs(client, "UnstoppableMango", "tdl", "v0.0.29")

		info, err := r.Stat("tdl-linux-amd64.tar.gz")

		Expect(err).NotTo(HaveOccurred())
		Expect(info.IsDir()).To(BeFalseBecause("release assets are files"))
		Expect(info.Name()).To(Equal("tdl-linux-amd64.tar.gz"))
	})

	It("should download an asset", func() {
		r := asset.NewFs(client, "UnstoppableMango", "tdl", "v0.0.29")

		file, err := r.Open("tdl-linux-amd64.tar.gz")

		Expect(err).NotTo(HaveOccurred())
		Expect(file.Name()).To(Equal("tdl-linux-amd64.tar.gz"))
		data, err := io.ReadAll(file)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(HaveLen(49_388_058))
	})
})
