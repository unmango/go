<!-- markdownlint-disable-file MD010 -->

# Go Utils

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/unmango/go/ci.yml)
![GitHub branch check runs](https://img.shields.io/github/check-runs/unmango/go/main)
![Codecov](https://img.shields.io/codecov/c/github/unmango/go)
![GitHub Release](https://img.shields.io/github/v/release/unmango/go)
![GitHub Release Date](https://img.shields.io/github/release-date/unmango/go)

Dumping ground for Go packages.
Some experiments with functional programming in packages like `iter`, `result`, and `rx`.
Some moderately useful abstractions in `codec` and `os`

## iter

The `iter` package builds on the standard `iter` package.
It re-exports `Seq` and `Seq2` for convenience and adds `Seq3`.
It adds sequence creation functions such as `Empty`, `Singleton`, and `Values`.

```go
var seq Seq[int] = iter.Empty[int]()
var seq Seq[int] = iter.Singleton(69)
var seq Seq[int] = iter.Values(69, 420)
```

It also adds various sequence operations such as `Map`, `Fold`, and `Filter`.

```go
var seq Seq[int] = iter.Values(69, 420)

// [70, 421]
mapped := iter.Map(seq, func(i int) int {
	return i + 1
})

// 489
sum := iter.Fold(seq, func(acc, i int) {
	return acc + i
}, 0)

// [69]
filtered := iter.Filter(seq, func(i int) bool {
	return i != 420
})
```

## maps

Primarily re-exports functions and types for convenience.

```go
func Test(seq iter.Seq2[string, int]) {
	var m map[string]int = maps.Collect(seq)
}
```

The `maps` package also adds `AppendSeq` for appending a `map` to a `Seq2`.

```go
seq := maps.All(map[string]string{"foo": "bar"})

// {"foo": "bar", "bin": "baz"}
seq = maps.AppendSeq(seq, map[string]string{"bin": "baz"})
```

## either

The `either` package adds the `Either` type representing exactly one of two possible values.
It also adds various operations such as `Map` and `Bind`.

```go
func main() {
	var m either.Either[int, string] = either.Left[int, string](420)

	m = either.MapLeft(m, func(x int) int {
		return x+1
	})

	// 421
	v, _ := m()
}
```

## either/result

The `result` pakcage adds the `Result` type representing either success or error.
It also adds various result operations such as `Map` and `Bind`.

```go
func main() {
	var r result.Result[int] = result.Ok(420)

	r = result.Map(r, func(x int) int {
		return x+1
	})

	// 421
	v, _ := r()
}
```

## either/maybe

The `maybe` pakcage adds the `Maybe` type representing either some value, or none.
It also adds various operations such as `Map` and `Bind`.

```go
func main() {
	var m maybe.Maybe[int] = maybe.Some(420)

	m = maybe.Map(m, func(x int) int {
		return x+1
	})

	// 421
	v, _ := m()
}
```

## slices

The `slices` package re-exports functions and types from the standard `slices` package for convenience.

## rx

The `rx` package attempts to implement the observable and signal patterns for reactive programming in go.
Both `observable` and `signal` should be considered a 🚧 work in progress 🚧, but the `observable` package is generally usable.

```go
var obs rx.Observable[int] = subject.New[int]()

sub := obs.Subscribe(observer.Lift(func(i int) {
	fmt.Println(i)
}))
defer sub()

obs.OnNext(69)
obs.OnComplete()
```

## References

A lot of this is inspired by the works of others, be sure to check out these repos as well.
They're much more interesting than anything you'll find here.

<https://github.com/fogfish/golem>

- [A Guide to Pure Combinators in Golang](https://medium.com/@dmkolesnikov/a-guide-to-pure-type-combinators-in-golang-or-how-to-stop-worrying-and-love-the-functional-e14f7f8cf35c)

<https://github.com/IBM/fp-go>
