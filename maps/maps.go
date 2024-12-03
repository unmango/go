package maps

import "github.com/unmango/go/iter"

func AppendSeq[M ~map[K]V, K comparable, V any](seq iter.Seq2[K, V], m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cont := true
		seq(func(k K, v V) bool {
			cont = yield(k, v)
			return cont
		})

		if !cont {
			return
		}

		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}
