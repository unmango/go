package make_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/make"
	"github.com/unmango/go/make/token"
)

var Identifiers = []TableEntry{
	Entry(nil, "ident"),
	Entry(nil, "./file/path"),
	Entry(nil, "/abs/path"),
	Entry(nil, "foo_bar"),
	Entry(nil, "foo-bar"),
	Entry(nil, "foo123"),
	Entry(nil, "123"),
	Entry(nil, "_foo"),
	Entry(nil, "foo_"),
	Entry(nil, "ident\n"),
	Entry(nil, "./file/path\n"),
	Entry(nil, "/abs/path\n"),
	Entry(nil, "foo_bar\n"),
	Entry(nil, "foo-bar\n"),
	Entry(nil, "foo123\n"),
	Entry(nil, "123\n"),
	Entry(nil, "_foo\n"),
	FEntry(nil, "foo_\n"),
}

var _ = Describe("Scanner", func() {
	DescribeTable("Scan identifier", Identifiers,
		func(input string) {
			buf := bytes.NewBufferString(input)
			s := make.NewScanner(buf, nil)

			Expect(s.Scan()).To(BeTrueBecause("more to scan"))
			Expect(s.Token()).To(Equal(token.IDENT))
			Expect(s.Literal()).To(Equal(strings.TrimSpace(input)))
			Expect(s.Scan()).To(BeTrueBecause("more to scan"))
			Expect(s.Token()).To(Equal(token.EOF))
		},
	)

	DescribeTable("Scan ident followed by token",
		Entry(nil, "ident $", Pending),
		Entry(nil, "ident:"),
		Entry(nil, "ident :"),
		Entry(nil, "ident =", Pending),
		Entry(nil, "ident :=", Pending),
		Entry(nil, "ident ::=", Pending),
		Entry(nil, "ident :::=", Pending),
		Entry(nil, "ident ?=", Pending),
		Entry(nil, "ident !=", Pending),
		Entry(nil, "ident (", Pending),
		Entry(nil, "ident )", Pending),
		Entry(nil, "ident {", Pending),
		Entry(nil, "ident }", Pending),
		Entry(nil, "ident ,", Pending),
		Entry(nil, "ident\n\t"),
		func(input string) {
			buf := bytes.NewBufferString(input)
			s := make.NewScanner(buf, nil)

			cont := s.Scan()

			Expect(s.Token()).To(Equal(token.IDENT))
			Expect(s.Literal()).To(Equal("ident"))
			Expect(cont).To(BeTrue())
		},
	)
})
