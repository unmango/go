package make

import (
	"io"

	"github.com/charmbracelet/log"
)

func Parse(r io.Reader) (*Makefile, error) {
	m := &Makefile{
		Rules: make([]Rule, 0),
	}

	s := NewScanner(r)
	for s.Scan() {
		switch tok := s.Token().(type) {
		case Rule:
			m.Rules = append(m.Rules, tok)
		default:
			log.Info("unexpected token", "tok", tok)
		}
	}

	if s.Err() != nil {
		return nil, s.Err()
	} else {
		return m, nil
	}
}
