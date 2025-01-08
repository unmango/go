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

	Describe("ScanTokens2", func() {
		DescribeTable("Scanner",
			Entry("target",
				"target:", []string{"target", ":"},
			),
			Entry("target with a separating space",
				"target :", []string{"target", ":"},
			),
			Entry("multiple targets",
				"target target2:", []string{"target", "target2", ":"},
			),
			Entry("multiple targets with a separating space",
				"target target2 :", []string{"target", "target2", ":"},
			),
			Entry("target with a trailing newline",
				"target:\n", []string{"target", ":", "\n"},
			),
			Entry("target with a prereq",
				"target: prereq", []string{"target", ":", "prereq"},
			),
			Entry("target with a prereq and trailing newline",
				"target: prereq\n", []string{"target", ":", "prereq", "\n"},
			),
			Entry("target with multiple prereqs",
				"target: prereq prereq2", []string{"target", ":", "prereq", "prereq2"},
			),
			Entry("target with a recipe",
				"target:\n\trecipe", []string{"target", ":", "\n", "\t", "recipe"},
			),
			Entry("target with a recipe and trailing newline",
				"target:\n\trecipe\n", []string{"target", ":", "\n", "\t", "recipe", "\n"},
			),
			Entry("target with multiple recipes",
				"target:\n\trecipe\n\trecipe2",
				[]string{"target", ":", "\n", "\t", "recipe", "\n", "\t", "recipe2"},
			),
			Entry("comment",
				"# comment", []string{"#", "comment"},
			),
			Entry("comment with multiple words",
				"# comment word", []string{"#", "comment", "word"},
			),
			Entry("comment with a trailing newline",
				"# comment\n", []string{"#", "comment", "\n"},
			),
			Entry("target with a comment",
				"target: # comment", []string{"target", ":", "#", "comment"},
			),
			Entry("directive",
				"define TEST", []string{"define", "TEST"},
			),
			Entry("prefixed include directive",
				"-include foo.mk", []string{"-include", "foo.mk"},
			),
			Entry("variable",
				"VAR := test", []string{"VAR", ":=", "test"},
			),
			Entry("variable with a trailing newline",
				"VAR := test\n", []string{"VAR", ":=", "test", "\n"},
			),
			func(text string, expected []string) {
				buf := bytes.NewBufferString(text)
				s := bufio.NewScanner(buf)
				s.Split(make.ScanTokens2)

				tokens := []string{}
				for s.Scan() {
					tokens = append(tokens, s.Text())
				}
				Expect(s.Err()).NotTo(HaveOccurred())
				Expect(tokens).To(Equal(expected))
			},
		)
	})
})
