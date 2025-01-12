package make

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/unmango/go/make/token"
)

type Scanner struct {
	s    *bufio.Scanner
	file *token.File

	tok token.Token
	lit string

	done       bool
	offset     int
	rdOffset   int
	lineOffset int
	nlPos      token.Pos
}

func NewScanner(r io.Reader, file *token.File) *Scanner {
	s := &Scanner{
		s:    bufio.NewScanner(r),
		file: file,
	}
	s.s.Split(ScanTokens)
	s.next()

	return s
}

func (s Scanner) Err() error {
	return s.s.Err()
}

func (s Scanner) Token() token.Token {
	return s.tok
}

func (s Scanner) Literal() string {
	return s.lit
}

func (s Scanner) Pos() token.Pos {
	return s.file.Pos(s.offset)
}

func (s *Scanner) Scan() bool {
	switch txt := s.s.Text(); {
	case token.IsIdentifier(txt):
		s.lit = txt
		if len(txt) > 1 {
			s.tok = token.Lookup(txt)
		} else {
			s.tok = token.IDENT
		}
	default:
		s.next()
		switch txt {
		case "=":
			s.tok = token.RECURSIVE_ASSIGN
		case ":=":
			s.tok = token.SIMPLE_ASSIGN
		case "::=":
			s.tok = token.POSIX_ASSIGN
		case ":::=":
			s.tok = token.IMMEDIATE_ASSIGN
		case "?=":
			s.tok = token.IFNDEF_ASSIGN
		case "!=":
			s.tok = token.SHELL_ASSIGN
		case ",":
			s.tok = token.COMMA
		case "\n":
			s.tok = token.NEWLINE
		case "\t":
			s.tok = token.TAB
		case "(":
			s.tok = token.LPAREN
		case ")":
			s.tok = token.RPAREN
		case "{":
			s.tok = token.LBRACE
		case "}":
			s.tok = token.RBRACE
		case "$":
			s.tok = token.DOLLAR
		case ":":
			s.tok = token.COLON
		case "#":
			s.lit = s.scanComment()
			s.tok = token.COMMENT
		default:
			s.tok = token.UNSUPPORTED
			s.lit = txt
		}
	}

	return !s.done
}

func (s *Scanner) next() {
	fmt.Println(s.done)
	if s.done {
		s.offset += len(s.s.Bytes())
		if s.isNewline() {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		s.tok = token.EOF
	} else {
		s.offset = s.rdOffset
		if s.isNewline() {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}

		s.rdOffset += len(s.s.Bytes())
		s.done = !s.s.Scan()
	}
}

func (s Scanner) isNewline() bool {
	return bytes.ContainsRune(s.s.Bytes(), '\n')
}

func (s *Scanner) skipWhitespace() {
	for bytes.ContainsAny(s.s.Bytes(), " \t\n\r") {
		s.next()
	}
}

func (s *Scanner) scanComment() string {
	b := &strings.Builder{}
	s.next() // Skip #
	for !s.done && !s.isNewline() {
		b.Write(s.s.Bytes())
		s.next()
	}
	if s.isNewline() {
		s.next() // Skip trailing \n
	}

	return b.String()
}
