package make

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

type Token string

const (
	CommentToken     Token = "comment"
	RuleToken        Token = "rule"
	UnsupportedToken Token = "unsupported"
	WhitespaceToken  Token = "whitespace"
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

func (s *Scanner) Token() interface{} {
	switch s.Type() {
	case RuleToken:
		lines := strings.Split(s.s.Text(), "\n")
		t, p, ok := strings.Cut(lines[0], ":")
		if !ok {
			panic("invalid rule")
		}

		recipe := []string{}
		if len(lines) > 1 {
			recipe = lines[1:]
		}

		return Rule{
			Target:  strings.Fields(t),
			PreReqs: strings.Fields(p),
			Recipe:  recipe,
		}
	case WhitespaceToken:
		return s.s.Text()
	case CommentToken:
		return s.s.Text()
	default:
		return nil
	}
}

func NewScanner(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanRules)

	return &Scanner{s: scanner}
}

func ScanRules(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if bytes.HasPrefix(data, []byte("\n")) {
		i := bytes.LastIndexFunc(data, unicode.IsSpace) + 1
		return i, data[:i], nil
	}

	lines := bytes.Split(data, []byte("\n"))
	if len(lines) == 0 {
		return 0, nil, nil
	}
	if bytes.ContainsRune(lines[0], ':') {
		advance = len(lines[0])
		for _, line := range lines[1:] {
			if bytes.HasPrefix(line, []byte("\t")) {
				advance += len(line) + 1
			} else {
				break
			}
		}

		return advance, data[:advance], nil
	}

	return 0, nil, nil
}
