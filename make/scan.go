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
		// Skip for now, I'll add support later
		fallthrough
	case '\t': // We're looking at a recipe
		// Consume until we find a comment or the end of the line
		if i := bytes.IndexFunc(data, LineOrComment); i > 0 {
			if data[i] == '#' { // If a comment is next we're done
				return i, data[:i], nil
			}
		}

		fallthrough // Otherwise consume the rest of the line
	case '#': // We're looking at a comment
		// The rest of the line should be a part of the comment
		if i := bytes.IndexRune(data, '\n'); i > 0 {
			i++ // Include the '\n'
			return i, data[:i], nil
		} else if atEOF {
			return len(data), data, nil
		} else {
			return 0, nil, nil
		}
	}

	if i := bytes.IndexRune(data, ':'); i > 0 {
		if i+1 < len(data) && data[i+1] == '=' { // We're looking at a variable
			// Consume until we find a comment or the end of the line
			for i = i + 1; i < len(data); i++ {
				if data[i] == '\n' {
					i++ // Include the '\n'
					break
				}
				if data[i] == '#' {
					break
				}
			}
		} else { // We're looking at a target
			// Eat the remaining space
			for i = i + 1; i < len(data); i++ {
				if data[i] == ' ' || data[i] == '\n' {
					continue
				} else {
					break
				}
			}
		}

		token = bytes.TrimRightFunc(data[:i], func(r rune) bool {
			return r != '\n' && unicode.IsSpace(r)
		})

		return i, token, nil
	}

	if atEOF {
		return len(data), data, nil
	} else {
		return 0, nil, nil
	}
}

func ScanTokens2(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
