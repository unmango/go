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

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Look for hints to avoid looping
	switch data[0] {
	case '$':
		// Skip the line for now, we'll add support later
		fallthrough
	case '#': // We're looking at a comment
		// The rest of the line should be a part of the comment
		if i := bytes.IndexRune(data, '\n'); i > 0 {
			return i, data[:i], nil
		} else {
			return 0, nil, nil
		}
	}

	// Are we looking at a target
	if i := bytes.IndexRune(data, ':'); i > 0 {
		if i+1 < len(data) && unicode.IsSpace(rune(data[i+1])) {
			// Increment to include the trailing whitespace
			i++
		}
		if s := bytes.IndexRune(data, ' '); s > 0 && s < i {
			// There are mutliple targets, take the first one
			i = s
		}

		// Add 1 to include the ':'
		return i + 1, data[:i+1], nil
	}

	return 0, nil, nil
}
