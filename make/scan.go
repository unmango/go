package make

import (
	"bytes"
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
		if len(data) > 1 && data[1] == '\t' {
			return 2, data[:2], nil
		}

		fallthrough
	case '\t':
		return 1, data[:1], nil
	}

	if i := bytes.IndexAny(data, ":\n\t "); i > 0 {
		switch data[i] {
		case ' ':
			return i + 1, data[:i], nil
		case ':', '\n', '\t':
			return i, data[:i], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	} else {
		return 0, nil, nil
	}
}
