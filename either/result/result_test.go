package result_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/go/either/result"
)

var _ = Describe("Result", func() {
	Describe("Ok", func() {
		It("should return the value with no error", func() {
			r := result.Ok(42)

			v, err := r()

			Expect(v).To(Equal(42))
			Expect(err).To(BeNil())
		})
	})

	Describe("Error", func() {
		It("should return zero value with error", func() {
			testErr := errors.New("test error")
			r := result.Error[int](testErr)

			v, err := r()

			Expect(v).To(Equal(0))
			Expect(err).To(BeIdenticalTo(testErr))
		})
	})

	Describe("ErrorString", func() {
		It("should return zero value with string error", func() {
			r := result.ErrorString[int]("test error")

			v, err := r()

			Expect(v).To(Equal(0))
			Expect(err).To(MatchError("test error"))
		})
	})

	Describe("Errorf", func() {
		It("should return zero value with formatted error", func() {
			r := result.Errorf[int]("error: %s %d", "test", 42)

			v, err := r()

			Expect(v).To(Equal(0))
			Expect(err).To(MatchError("error: test 42"))
		})
	})

	Describe("From", func() {
		It("should return the value with no error", func() {
			r := result.From(42, nil)

			v, err := r()

			Expect(v).To(Equal(42))
			Expect(err).To(BeNil())
		})

		It("should return the value with error", func() {
			testErr := errors.New("test error")
			r := result.From(42, testErr)

			v, err := r()

			Expect(v).To(Equal(42))
			Expect(err).To(BeIdenticalTo(testErr))
		})
	})

	Describe("Map", func() {
		It("should map the ok value", func() {
			r := result.Ok(42)

			mapped := result.Map(r, func(v int) int {
				return v * 2
			})

			v, err := mapped()
			Expect(v).To(Equal(84))
			Expect(err).To(BeNil())
		})

		It("should not map the error value", func() {
			testErr := errors.New("test error")
			r := result.Error[int](testErr)

			mapped := result.Map(r, func(v int) int {
				return v * 2
			})

			v, err := mapped()
			Expect(v).To(Equal(0))
			Expect(err).To(BeIdenticalTo(testErr))
		})
	})

	Describe("Bind", func() {
		It("should bind the ok value", func() {
			r := result.Ok(42)

			bound := result.Bind(r, func(v int) result.Result[int] {
				return result.Ok(v * 2)
			})

			v, err := bound()
			Expect(v).To(Equal(84))
			Expect(err).To(BeNil())
		})

		It("should not bind the error value", func() {
			testErr := errors.New("test error")
			r := result.Error[int](testErr)

			bound := result.Bind(r, func(v int) result.Result[int] {
				return result.Ok(v * 2)
			})

			v, err := bound()
			Expect(v).To(Equal(0))
			Expect(err).To(BeIdenticalTo(testErr))
		})

		It("should propagate error from bind function", func() {
			r := result.Ok(42)
			bindErr := errors.New("bind error")

			bound := result.Bind(r, func(v int) result.Result[int] {
				return result.Error[int](bindErr)
			})

			v, err := bound()
			Expect(v).To(Equal(0))
			Expect(err).To(BeIdenticalTo(bindErr))
		})
	})

	Context("Result2", func() {
		Describe("Ok2", func() {
			It("should return both values with no error", func() {
				r := result.Ok2(42, "test")

				v1, v2, err := r()

				Expect(v1).To(Equal(42))
				Expect(v2).To(Equal("test"))
				Expect(err).To(BeNil())
			})
		})

		Describe("Error2", func() {
			It("should return zero values with error", func() {
				testErr := errors.New("test error")
				r := result.Error2[int, string](testErr)

				v1, v2, err := r()

				Expect(v1).To(Equal(0))
				Expect(v2).To(Equal(""))
				Expect(err).To(BeIdenticalTo(testErr))
			})
		})

		Describe("From2", func() {
			It("should return both values with no error", func() {
				r := result.From2(42, "test", nil)

				v1, v2, err := r()

				Expect(v1).To(Equal(42))
				Expect(v2).To(Equal("test"))
				Expect(err).To(BeNil())
			})

			It("should return both values with error", func() {
				testErr := errors.New("test error")
				r := result.From2(42, "test", testErr)

				v1, v2, err := r()

				Expect(v1).To(Equal(42))
				Expect(v2).To(Equal("test"))
				Expect(err).To(BeIdenticalTo(testErr))
			})
		})

		Describe("Map2", func() {
			It("should map both ok values", func() {
				r := result.Ok2(42, "test")

				mapped := result.Map2(r, func(v1 int, v2 string) (int, string) {
					return v1 * 2, v2 + "!"
				})

				v1, v2, err := mapped()
				Expect(v1).To(Equal(84))
				Expect(v2).To(Equal("test!"))
				Expect(err).To(BeNil())
			})

			It("should not map when there is an error", func() {
				testErr := errors.New("test error")
				r := result.Error2[int, string](testErr)

				mapped := result.Map2(r, func(v1 int, v2 string) (int, string) {
					return v1 * 2, v2 + "!"
				})

				v1, v2, err := mapped()
				Expect(v1).To(Equal(0))
				Expect(v2).To(Equal(""))
				Expect(err).To(BeIdenticalTo(testErr))
			})
		})

		Describe("Bind2", func() {
			It("should bind both ok values", func() {
				r := result.Ok2(42, "test")

				bound := result.Bind2(r, func(v1 int, v2 string) result.Result2[int, string] {
					return result.Ok2(v1*2, v2+"!")
				})

				v1, v2, err := bound()
				Expect(v1).To(Equal(84))
				Expect(v2).To(Equal("test!"))
				Expect(err).To(BeNil())
			})

			It("should not bind when there is an error", func() {
				testErr := errors.New("test error")
				r := result.Error2[int, string](testErr)

				bound := result.Bind2(r, func(v1 int, v2 string) result.Result2[int, string] {
					return result.Ok2(v1*2, v2+"!")
				})

				v1, v2, err := bound()
				Expect(v1).To(Equal(0))
				Expect(v2).To(Equal(""))
				Expect(err).To(BeIdenticalTo(testErr))
			})

			It("should propagate error from bind function", func() {
				r := result.Ok2(42, "test")
				bindErr := errors.New("bind error")

				bound := result.Bind2(r, func(v1 int, v2 string) result.Result2[int, string] {
					return result.Error2[int, string](bindErr)
				})

				v1, v2, err := bound()
				Expect(v1).To(Equal(0))
				Expect(v2).To(Equal(""))
				Expect(err).To(BeIdenticalTo(bindErr))
			})
		})
	})
})
