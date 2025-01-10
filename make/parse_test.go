package make_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/make"
)

var _ = Describe("Parse", func() {
	It("should parse a target", func() {
		buf := bytes.NewBufferString(`target:`)

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
	})

	It("should parse a prereq", func() {
		buf := bytes.NewBufferString("target: prereq\n")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{"prereq"},
			Recipe:  []string{},
		}))
	})

	It("should parse multiple prereqs", func() {
		buf := bytes.NewBufferString("target: prereq prereq2\n")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{"prereq", "prereq2"},
			Recipe:  []string{},
		}))
	})

	It("should parse a rule with multiple targets", func() {
		buf := bytes.NewBufferString(`target target2:`)

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target", "target2"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
	})

	It("should parse multiple targets", func() {
		buf := bytes.NewBufferString("target:\ntarget2:")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(
			make.Rule{
				Target:  []string{"target"},
				PreReqs: []string{},
				Recipe:  []string{},
			},
			make.Rule{
				Target:  []string{"target2"},
				PreReqs: []string{},
				Recipe:  []string{},
			},
		))
	})

	It("should ignore leading newlines", func() {
		buf := bytes.NewBufferString("\ntarget:")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
	})

	It("should ignore separating newlines", func() {
		buf := bytes.NewBufferString("target:\n\ntarget2:")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(
			make.Rule{
				Target:  []string{"target"},
				PreReqs: []string{},
				Recipe:  []string{},
			},
			make.Rule{
				Target:  []string{"target2"},
				PreReqs: []string{},
				Recipe:  []string{},
			},
		))
	})

	It("should ignore trailing newlines", func() {
		buf := bytes.NewBufferString("target:\n")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{},
		}))
	})

	It("should parse recipes", func() {
		buf := bytes.NewBufferString("target:\n\trecipe\n")

		m, err := make.Parse(buf)

		Expect(err).NotTo(HaveOccurred())
		Expect(m.Rules).To(ConsistOf(make.Rule{
			Target:  []string{"target"},
			PreReqs: []string{},
			Recipe:  []string{"recipe"},
		}))
	})
})
