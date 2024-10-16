package seqs_test

import (
	"errors"
	"testing/quick"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/iter"
	"github.com/unmango/go/iter/seqs"
	"github.com/unmango/go/result"
)

var _ = Describe("Seq3", func() {
	Context("Filter3R", func() {
		It("should exlude err results", func() {
			s := iter.Singleton3(0, 0, result.Err[int](errors.New("test err")))

			result := seqs.Filter3R(s)

			sentinel := false
			result(func(int, int, int) bool {
				sentinel = true
				return true
			})

			Expect(sentinel).To(BeFalseBecause("the sequence should be empty"))
		})

		It("should include success results", func() {
			f := func(a0, b0, c0 int) bool {
				s := iter.Singleton3(a0, b0, result.Ok(c0))

				result := seqs.Filter3R(s)

				var a, b, c int
				result(func(a1, b1, c1 int) bool {
					a, b, c = a1, b1, c1
					return true
				})

				return a == a0 && b == b0 && c == c0
			}

			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("Map3", func() {
		It("should map", func() {
			f := func(a0, b0, c0 int) bool {
				s := iter.Singleton3(a0, b0, c0)

				result := seqs.Map3(s, func(a, b, c int) (int, int, int) {
					return a + 1, b + 1, c + 1
				})

				var a, b, c int
				result(func(a1, b1, c1 int) bool {
					a, b, c = a1, b1, c1
					return true
				})

				return a == a0+1 && b == b0+1 && c == c0+1
			}

			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})
})
