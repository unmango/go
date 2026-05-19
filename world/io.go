package world

import "github.com/unmango/go/world/os"

var System = system{os.System}

type IO interface {
	Os() Os
}

type system struct {
	os Os
}

func (s system) Os() Os {
	return s.os
}
