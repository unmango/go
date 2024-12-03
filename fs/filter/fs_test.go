package filter_test

import (
	"os"
	"syscall"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/go/fs/filter"
)

var _ = Describe("Fs", func() {
	It("should not allow chmod-ing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		err = filtered.Chmod("test.txt", os.ModeAppend)

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should not allow chown-ing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		err = filtered.Chown("test.txt", 1001, 1001)

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should not allow chtimes-ing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		err = filtered.Chtimes("test.txt", time.Now(), time.Now())

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should not allow creating a filtered file", func() {
		fs := afero.NewMemMapFs()
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		_, err := filtered.Create("test.txt")

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should include the source filesystem name", func() {
		fs := filter.NewFs(afero.NewMemMapFs(), func(s string) bool {
			return true
		})

		Expect(fs.Name()).To(Equal("Filter: MemMapFS"))
	})

	It("should not allow opening a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		_, err = filtered.Open("test.txt")

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should not allow open-file-ing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		_, err = filtered.OpenFile("test.txt", 69, os.ModeAppend)

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should not allow removing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		err = filtered.Remove("test.txt")

		Expect(err).To(MatchError(syscall.ENOENT))
	})

	It("should allow removing a directory", func() {
		fs := afero.NewMemMapFs()
		err := fs.Mkdir("test", os.ModeDir)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		err = filtered.Remove("test")

		Expect(err).NotTo(HaveOccurred())
	})

	It("should not allow stat-ing a filtered file", func() {
		fs := afero.NewMemMapFs()
		err := afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
		filtered := filter.NewFs(fs, func(name string) bool {
			return name == "not-test.txt"
		})

		_, err = filtered.Stat("test.txt")

		Expect(err).To(MatchError(syscall.ENOENT))
	})
})
