package slices

import (
	"slices"

	"github.com/unmango/go/iter"
)

func Collect[E any](s iter.Seq[E]) []E {
	return slices.Collect(iter.D(s))
}
