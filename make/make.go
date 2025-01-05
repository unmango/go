package make

import (
	"io"
)

type Rule struct {
	Target  []string
	PreReqs []string
	Recipe  []string
}

type Makefile struct {
	Rules []Rule
}

func Parse(r io.Reader) (*Makefile, error) {
	m := &Makefile{
		Rules: make([]Rule, 0),
	}

	s := NewScanner(r)
	for s.Scan() {
		token := s.Token()
		if token == nil {
			continue
		}

		switch n := token.(type) {
		case Rule:
			m.Rules = append(m.Rules, n)
		}
	}

	if s.Err() != nil {
		return nil, s.Err()
	} else {
		return m, nil
	}
}
