package make

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Token string

const (
	CommentToken      Token = "comment"
	DirectiveToken    Token = "directive"
	PreRequisiteToken Token = "prerequisite"
	RecipeToken       Token = "recipe"
	RuleToken         Token = "rule"
	TargetToken       Token = "target"
	UnsupportedToken  Token = "unsupported"
	VariableToken     Token = "variable"
	WhitespaceToken   Token = "whitespace"
)

type Scanner struct {
	s *bufio.Scanner
}

func (s *Scanner) Split(split bufio.SplitFunc) {
	s.s.Split(split)
}

func (s *Scanner) Type() Token {
	text := s.s.Text()
	if strings.TrimSpace(text) == "" {
		return WhitespaceToken
	}
	if strings.HasPrefix(text, "#") {
		return CommentToken
	}

	lines := strings.Split(s.s.Text(), "\n")
	if len(lines) == 0 {
		panic("newline delimeted text was empty")
	}

	switch {
	case strings.ContainsRune(lines[0], ':'):
		return RuleToken
	default:
		return UnsupportedToken
	}
}

func (s *Scanner) Err() error {
	return s.s.Err()
}

func (s *Scanner) Scan() bool {
	return s.s.Scan()
}

func NewScanner(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanTokens)

	return &Scanner{s: scanner}
}

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	switch data[0] {
	case ' ':
		return 1, nil, nil
	case ':':
		if len(data) > 1 && data[1] == '=' {
			return 2, data[:2], nil
		}

		fallthrough
	case '#':
		if len(data) > 1 && data[1] == ' ' {
			return 2, data[:1], nil
		}

		fallthrough
	case '\n':
		fallthrough
	case '\t':
		return 1, data[:1], nil
	}

	if i := bytes.IndexAny(data, ":\n\t "); i > 0 {
		switch data[i] {
		case ' ':
			return i + 1, data[:i], nil
		case ':':
			return i, data[:i], nil
		case '\n':
			fallthrough
		case '\t':
			return i, data[:i], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	} else {
		return 0, nil, nil
	}
}

func LineOrComment(r rune) bool {
	return r == '\n' || r == '#'
}
