package docker_test

import (
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/fs/docker"
)

var _ = Describe("Fs", func() {
	It("should list directories", func() {
		fs := docker.NewFs(testclient, ctr.GetContainerID())
		dir, err := fs.Open("/")
		Expect(err).NotTo(HaveOccurred())

		infos, err := dir.Readdir(69)

		Expect(err).NotTo(HaveOccurred())
		names := make([]string, len(infos))
		for i, f := range infos {
			names[i] = f.Name()
		}
		Expect(names).To(ContainElements("root", "var", "bin"))
	})

	It("should read file contents", func() {
		fs := docker.NewFs(testclient, ctr.GetContainerID())
		file, err := fs.Create("test-read.txt")
		Expect(err).NotTo(HaveOccurred())
		_, err = io.WriteString(file, "bleh")
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Close()).To(Succeed())

		file, err = fs.Open("test-read.txt")

		Expect(err).NotTo(HaveOccurred())
		data, err := io.ReadAll(file)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(data)).To(Equal("bleh"))
	})

	Describe("Create", func() {
		It("should create a file", func() {
			fsys := docker.NewFs(testclient, ctr.GetContainerID())

			file, err := fsys.Create("test.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(file).NotTo(BeNil())
		})

		It("should create a writable file", func() {
			fsys := docker.NewFs(testclient, ctr.GetContainerID())

			file, err := fsys.Create("writable.txt")

			Expect(err).NotTo(HaveOccurred())
			_, err = file.WriteString("blahblahblah")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
