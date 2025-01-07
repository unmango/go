package make_test

import (
	"bufio"
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/make"
)

var _ = Describe("Scan", func() {
	Describe("ScanRules", func() {
		It("should split a rule with a single target", func() {
			buf := []byte("target:")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(7))
		})

		It("should split a rule with a single target and a prereq", func() {
			buf := []byte("target: prereq")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(14))
		})

		It("should split a rule with a single target and multiple prereqs", func() {
			buf := []byte("target: prereq prereq2")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(22))
		})

		It("should split a rule with multiple targets", func() {
			buf := []byte("target target2:")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(15))
		})

		It("should split a rule with multiple targets and a prereq", func() {
			buf := []byte("target target2: prereq")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(22))
		})

		It("should split a rule with multiple targets and multiple prereqs", func() {
			buf := []byte("target target2: prereq prereq2")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(30))
		})

		It("should split a rule with a recipe", func() {
			buf := []byte("target:\n\trecipe")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(15))
		})

		It("should split a rule with a prereq and a recipe", func() {
			buf := []byte("target: prereq\n\trecipe")
			advance, token, err := make.ScanRules(buf, true)

			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal(buf))
			Expect(advance).To(Equal(22))
		})
	})

	Describe("ScanTokens", func() {
		It("should split a target", func() {
			buf := bytes.NewBufferString("target:")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:"))
		})

		It("should split a target with a separating space", func() {
			buf := bytes.NewBufferString("target :")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target :"))
		})

		It("should split multiple targets", func() {
			buf := bytes.NewBufferString("target target2:")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target target2:"))
		})

		It("should split multiple targets with a separating space", func() {
			buf := bytes.NewBufferString("target target2 :")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target target2 :"))
		})

		It("should split a target with a trailing newline", func() {
			buf := bytes.NewBufferString("target:\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:\n"))
		})

		It("should split a target with a prereq", func() {
			buf := bytes.NewBufferString("target: prereq")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq"))
		})

		It("should split a target with a prereq and trailing newline", func() {
			buf := bytes.NewBufferString("target: prereq\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq\n"))
		})

		It("should split a target with prereqs", func() {
			buf := bytes.NewBufferString("target: prereq prereq2")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq prereq2"))
		})

		It("should split a target with a recipe", func() {
			buf := bytes.NewBufferString("target:\n\trecipe")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\trecipe"))
		})

		It("should split a target with a recipe and trailing newline", func() {
			buf := bytes.NewBufferString("target:\n\trecipe\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\trecipe\n"))
		})

		It("should split a target with multiple recipes", func() {
			buf := bytes.NewBufferString("target:\n\trecipe\n\trecipe2")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\trecipe\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\trecipe2"))
		})

		It("should split a comment", func() {
			buf := bytes.NewBufferString("# comment")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("# comment"))
		})

		It("should split a comment with a trailing newline", func() {
			buf := bytes.NewBufferString("# comment\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("# comment\n"))
		})

		It("should split target with a comment", func() {
			buf := bytes.NewBufferString("target: # comment")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target:"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("# comment"))
		})

		It("should split a directive", func() {
			buf := bytes.NewBufferString("define TEST")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("define TEST"))
		})

		It("should split a prefixed include directive", func() {
			buf := bytes.NewBufferString("-include foo.mk")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("-include foo.mk"))
		})

		It("should split a variable", func() {
			buf := bytes.NewBufferString("VAR := test")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("VAR := test"))
		})

		It("should split a variable with a trailing newline", func() {
			buf := bytes.NewBufferString("VAR := test\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("VAR := test\n"))
		})
	})

	FDescribe("ScanTokens2", func() {
		It("should split a target", func() {
			buf := bytes.NewBufferString("target:")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
		})

		It("should split a target with a separating space", func() {
			buf := bytes.NewBufferString("target :")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
		})

		It("should split multiple targets", func() {
			buf := bytes.NewBufferString("target target2:")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target2"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
		})

		It("should split multiple targets with a separating space", func() {
			buf := bytes.NewBufferString("target target2 :")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target2"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
		})

		It("should split a target with a trailing newline", func() {
			buf := bytes.NewBufferString("target:\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
		})

		It("should split a target with a prereq", func() {
			buf := bytes.NewBufferString("target: prereq")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq"))
		})

		It("should split a target with a prereq and trailing newline", func() {
			buf := bytes.NewBufferString("target: prereq\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
		})

		It("should split a target with prereqs", func() {
			buf := bytes.NewBufferString("target: prereq prereq2")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("prereq2"))
		})

		It("should split a target with a recipe", func() {
			buf := bytes.NewBufferString("target:\n\trecipe")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\t"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("recipe"))
		})

		It("should split a target with a recipe and trailing newline", func() {
			buf := bytes.NewBufferString("target:\n\trecipe\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\t"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("recipe"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
		})

		It("should split a target with multiple recipes", func() {
			buf := bytes.NewBufferString("target:\n\trecipe\n\trecipe2")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\t"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("recipe"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\t"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("recipe2"))
		})

		It("should split a comment", func() {
			buf := bytes.NewBufferString("# comment")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("#"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("comment"))
		})

		It("should split a comment into words", func() {
			buf := bytes.NewBufferString("# comment word")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("#"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("comment"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("word"))
		})

		It("should split a comment with a trailing newline", func() {
			buf := bytes.NewBufferString("# comment\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("#"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("comment"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
		})

		It("should split target with a comment", func() {
			buf := bytes.NewBufferString("target: # comment")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("target"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("#"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("comment"))
		})

		It("should split a directive", func() {
			buf := bytes.NewBufferString("define TEST")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("define"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("TEST"))
		})

		It("should split a prefixed include directive", func() {
			buf := bytes.NewBufferString("-include foo.mk")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("-include"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("foo.mk"))
		})

		It("should split a variable", func() {
			buf := bytes.NewBufferString("VAR := test")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("VAR"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":="))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("test"))
		})

		It("should split a variable with a trailing newline", func() {
			buf := bytes.NewBufferString("VAR := test\n")
			s := bufio.NewScanner(buf)
			s.Split(make.ScanTokens2)

			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("VAR"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal(":="))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("test"))
			Expect(s.Scan()).To(BeTrue())
			Expect(s.Text()).To(Equal("\n"))
		})
	})
})
