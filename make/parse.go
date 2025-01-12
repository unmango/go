package make

import (
	"bufio"
	"io"
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

var emtpyRule = Rule{
	Target:  []string{},
	PreReqs: []string{},
	Recipe:  []string{},
}

type builder struct {
	m    Makefile
	buf  []string
	idx  int
	next Token
}

func newBuilder() *builder {
	return &builder{
		m:   Makefile{},
		idx: -1,
	}
}

func (b *builder) append(tok string) {
	b.buf = append(b.buf, tok)
}

func (b *builder) newRule() {
	b.m.Rules = append(b.m.Rules, emtpyRule)
	b.idx++
	b.start(TargetToken)
	b.end()
}

func (b *builder) start(t Token) {
	b.next = t
}

func (b *builder) end() {
	if b.buf == nil || len(b.buf) == 0 {
		return
	}

	switch b.next {
	case TargetToken:
		b.m.Rules[b.idx].Target = b.buf
	case PreRequisiteToken:
		b.m.Rules[b.idx].PreReqs = b.buf
	case RecipeToken:
		b.m.Rules[b.idx].Recipe = b.buf
	}

	b.buf = nil
}

func Parse(r io.Reader) (*Makefile, error) {
	s := bufio.NewScanner(r)
	s.Split(ScanTokens)

	b := newBuilder()

	for s.Scan() {
		switch tok := s.Text(); tok {
		case ":":
			b.newRule()
			b.start(PreRequisiteToken)
		case "\n\t":
			b.start(RecipeToken)
		case "\n":
			b.end()
		default:
			b.append(tok)
		}
	}

	if s.Err() != nil {
		return nil, s.Err()
	} else {
		return &b.m, nil
	}
}
