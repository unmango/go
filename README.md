<!-- markdownlint-disable-file MD010 -->

# Go Source code for UnMango

A collection of random Go packages.

## iter

Builds on the standard `iter` package.
Re-exports `Seq` and `Seq2` and adds `Seq3`.
Adds a few sequence creation functions such as `Empty` and `Singleton`.

```go
var seq Seq[int] = iter.Empty[int]()
var seq Seq[int] = iter.Singleton[int]()
```

## maps

Currently only adds `Collect` for creating a `map[]` from a `Seq2`.

```go
func Test(seq iter.Seq2[string, int]) {
	var m map[string]int = maps.Collect(seq)
}
```

## result

Adds the `Result` type representing either success or error.
Adds various result operations such as `Map` and `Bind`.

```go
func main() {
	var r Result[int] = func() (int, error) {
		return 420, nil
	}

	r = result.Map(r, func(x int) int {
		return x+1
	})
}
```

## seqs

Adds various sequence operations such as `Map`.

```go
func Test(seq iter.Seq[int]) {
	seqs.Map(seq, func(x int) int {
		return x+1
	})
}
```

## slices

Currently only re-exports `Collect`.

## Inspiration

I stand on the shoulders of giants.
A lot of this is inspired by the works of others, be sure to check out these repos as well.
(They're much smarter than me)

<https://github.com/fogfish/golem>

- [A Guide to Pure Combinators in Golang](https://medium.com/@dmkolesnikov/a-guide-to-pure-type-combinators-in-golang-or-how-to-stop-worrying-and-love-the-functional-e14f7f8cf35c)

<https://github.com/IBM/fp-go>
