package make_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/make"
)

var _ = Describe("Scan", func() {
	It("should scan a target", func() {
		buf := bytes.NewBufferString(`target:`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should scan prereqs", func() {
		buf := bytes.NewBufferString(`target: prereq`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{"prereq"},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should scan multiple prereqs", func() {
		buf := bytes.NewBufferString(`target: prereq prereq2`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{"prereq", "prereq2"},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should scan a rule with multiple targets", func() {
		buf := bytes.NewBufferString(`target target2:`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target", "target2"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should scan multiple targets", func() {
		buf := bytes.NewBufferString(`target:
target2:`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal("\n"))

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target2"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should ignore leading newlines", func() {
		buf := bytes.NewBufferString("\ntarget:")
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrueBecause("newline was scanned"))
		Expect(s.Token()).To(Equal("\n"))

		Expect(s.Scan()).To(BeTrueBecause("target was scanned"))
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should ignore separating newlines", func() {
		buf := bytes.NewBufferString(`target:

target2:`)
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal("\n\n"))

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target2"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

	It("should ignore trailing newlines", func() {
		buf := bytes.NewBufferString("target:\n")
		s := make.NewScanner(buf)

		Expect(s.Scan()).To(BeTrue())
		Expect(s.Token()).To(Equal(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))

		Expect(s.Scan()).To(BeTrueBecause("newline was scanned"))
		Expect(s.Token()).To(Equal("\n"))
		Expect(s.Scan()).To(BeFalseBecause("scanner should be empty"))
	})

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
})
