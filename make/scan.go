package make

import (
	"bufio"
	"bytes"
	"io"
	"slices"
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

type SplitFunc func(data []string, atEOF bool) (advance int, token interface{}, err error)

const maxConsecutiveEmptyReads = 100

// Scanner provides a convenient interface for reading tokens from a Makefile.
//
// Deprecated: This implementation is a WIP and is not yet stable.
type Scanner struct {
	s          *bufio.Scanner
	split      SplitFunc
	token      interface{}
	buf        []string
	empties    int
	err        error
	scanCalled bool
	done       bool
}

func (s *Scanner) Split(split SplitFunc) {
	if s.scanCalled {
		panic("Split called after Scan")
	}

	s.split = split
}

func (s *Scanner) Token() interface{} {
	return s.token
}

func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) Scan() bool {
	// Nearly identical logic as bufio except [token]
	// is an interface{} and [buf] is a []string

	if s.done {
		return false
	}
	s.scanCalled = true

	for {
		if len(s.buf) > 0 || s.err != nil {
			advance, token, err := s.split(s.buf, s.err != nil || s.done)
			if err != nil {
				if err == bufio.ErrFinalToken {
					s.token = token
					s.done = true
					return token != nil
				} else {
					s.setErr(err)
					return false
				}
			}

			if !s.advance(advance) {
				return false
			}

			s.token = token
			if token != nil {
				if s.err == nil || advance > 0 {
					s.empties = 0
				} else {
					s.empties++
					if s.empties > maxConsecutiveEmptyReads {
						panic("make.Scan: too many empty tokens without progressing")
					}
				}

				return true
			}
		}

		if s.err != nil {
			s.buf = []string{}
			return false
		}

		if s.s.Scan() {
			s.buf = append(s.buf, s.s.Text())
		} else {
			s.done = true
		}
	}
}

func (s *Scanner) advance(n int) bool {
	if n < 0 {
		s.setErr(bufio.ErrNegativeAdvance)
		return false
	}
	if n > len(s.buf) {
		s.setErr(bufio.ErrAdvanceTooFar)
		return false
	}

	s.buf = s.buf[n:]
	return true
}

func (s *Scanner) setErr(err error) {
	if s.err == nil || s.err == io.EOF {
		s.err = err
	}
}

// NewScanner returns a new Scanner to read from r.
//
// Deprecated: This implementation is a WIP and is not yet stable.
func NewScanner(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanTokens)

	return &Scanner{
		s:     scanner,
		split: ScanRules,
	}
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

// ScanRules is a split function for a Scanner that returns make rule tokens.
//
// Deprecated: This implementation is a WIP and is not yet stable.
func ScanRules(data []string, atEOF bool) (advance int, token interface{}, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	var colonIdx, end int
	if colonIdx = slices.Index(data, ":"); colonIdx < 0 {
		return 0, nil, nil // We haven't reached a rule yet
	}
	if end = slices.Index(data[colonIdx:], "\n"); end < 0 && !atEOF {
		return 0, nil, nil // We haven't reached the end of the rule yet
	}
	// TODO: Recipe

	r := Rule{
		Target:  data[:colonIdx],
		PreReqs: data[colonIdx+1:],
		Recipe:  []string{},
	}

	return 0, r, nil // TODO
}
