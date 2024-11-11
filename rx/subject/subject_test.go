package subject_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/rx/subject"
)

type testSub struct {
	onComplete func()
	onError    func(error)
	onNext     func(int)
}

// OnComplete implements rx.Observer.
func (t *testSub) OnComplete() {
	t.onComplete()
}

// OnError implements rx.Observer.
func (t *testSub) OnError(err error) {
	t.onError(err)
}

// OnNext implements rx.Observer.
func (t *testSub) OnNext(x int) {
	t.onNext(x)
}

var _ = Describe("Subject", func() {
	It("should initialize", func() {
		s := subject.New[int]()

		Expect(s).NotTo(BeNil())
	})

	It("should report next", func() {
		var c, n, e bool
		s := subject.New[int]()

		s.Subscribe(&testSub{
			onComplete: func() { c = true },
			onError:    func(error) { e = true },
			onNext:     func(int) { n = true },
		})
		s.OnComplete()
		s.OnNext(69)
		s.OnError(nil)

		Expect(c).To(BeTrueBecause("complete was called"))
		Expect(n).To(BeTrueBecause("next was called"))
		Expect(e).To(BeTrueBecause("error was called"))
	})
})
