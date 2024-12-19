<!-- markdownlint-disable-file MD010 -->

# Go Source code for UnMango

A collection of random Go packages.

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
Due to Go not currently supporting generic type aliases, these functions adapt the standard `iter` seq to this module's `iter` package.

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

## result

The `result` pakcage adds the `Result` type representing either success or error.
It also adds various result operations such as `Map` and `Bind`.

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

## slices

The `slices` package re-exports functions and types from the standard `slices` package for convenience.
Due to Go not currently supporting generic type aliases, these functions adapt the standard `slices` seq to this module's `slices` package.

## fs

The `fs` packages expands on [`github.com/spf13/afero`](https://github.com/spf13/afero) by adding more `afero.Fs` implementations as well as various `afero.Fs` utility functions.

### context

The `context` package adds the `context.Fs` interface for filesystem implementations that accept a `context.Context` per operation.
It has a basic test suite and generally works but should be considered a 🚧 work in progress 🚧.

The `context` package re-exports various functions and types from the standard `context` packge for convenience.
Currently the creation functions focus on adapting external `context.Fs` implementations to an `afero.Fs` to be used with the other utility functions.

```go
var base context.Fs = mypkg.NewEffectfulFs()

fs := context.BackgroundFs(base)

var accessor context.AccessorFunc = func() context.Context {
	return context.Background()
}

// Equivalent to `context.BackgroundFs`
fs := context.NewFs(base, accessor)
```

The `context.AferoFs` interface is a union of `afero.Fs` and `context.Fs`, i.e. exposing both `fs.Create` and `fs.CreateContext`.
I'm not sure if this actually has any value but it exists.

The `context.Discard` function adapts an `afero.Fs` to a `context.AferoFs` by ignoring the `context.Context` argument.

```go
base := afero.NewMemMapFs()

var fs context.AferoFs = context.Discard(base)
```

### docker

The `docker` package adds a docker `afero.Fs` implementation for operating on the filesystem of a container.

```go
client := client.NewClientWithOpts(client.FromEnv)

fs := docker.NewFs(client, "my-container-id")
```

### filter

The `filter` package adds a filtering implementation of `afero.Fs` similar to `afero.RegExpFs` at accepts a predicate instead.

```go
base := afero.NewMemMapFs()

fs := filter.NewFs(base, func(path string) bool {
	return filepath.Ext(path) == ".go"
})
```

### github

The `github` package adds multiple implementations of `afero.Fs` for interacting with the GitHub API as if it were a filesystem.
In general it can turn a GitHub url into an `afero.Fs`.

```go
fs := github.NewFs(github.NewClient(nil))

file, _ := fs.Open("https://github.com/unmango")

// ["go", "thecluster", "pulumi-baremetal", ...]
file.Readdirnames(420)
```

### ignore

The `ignore` package adds a filtering `afero.Fs` that accepts a `.gitignore` file and ignores paths matched by it.

```go
base := afero.NewMemMapFs()

gitignore, _ := os.Open("path/to/my/.gitignore")

fs, _ := ignore.NewFsFromGitIgnoreReader(base, gitignore)
```

### testing

The `testing` package adds helper stubs for mocking filesystems in tests.

```go
fs := &testing.Fs{
	CreateFunc: func(name string) (afero.File, error) {
		return nil, errors.New("simulated error")
	}
}
```

### writer

The `writer` package adds a readonly `afero.Fs` implementation that dumps all file writes to the provided `io.Writer`.
Currently paths are ignored and there are no delimeters separating files.

```go
buf := &bytes.Buffer{}
fs := writer.NewFs(buf)

_ = afero.WriteFile(fs, "test.txt", []byte("testing"), os.ModePerm)
_ = afero.WriteFile(fs, "other.txt", []byte("blah"), os.ModePerm)

// "testingblah"
buf.String()
```

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

## cmd

The `cmd` package contains CLI utilities exposed by the repo.
Currently the only published tool is `devops` which assists with listing files for `make` targets.

```shell
devops list --go
# cmd/devops/main.go
# devops/cmd/list.go
# fp/constraint/constraints.go
# ...
```

```makefile
bin/my-tool: $(shell devops list --go --exclude-tests)
	go build -o $@
```

## Inspiration

I stand on the shoulders of giants.
A lot of this is inspired by the works of others, be sure to check out these repos as well.
(They're much smarter than me)

<https://github.com/fogfish/golem>

- [A Guide to Pure Combinators in Golang](https://medium.com/@dmkolesnikov/a-guide-to-pure-type-combinators-in-golang-or-how-to-stop-worrying-and-love-the-functional-e14f7f8cf35c)

<https://github.com/IBM/fp-go>
