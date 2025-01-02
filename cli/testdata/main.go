package main

import (
	"os"

	"github.com/unmango/go/cli"
)

func main() {
	var a []any
	for _, x := range os.Args[1:] {
		a = append(a, x)
	}

	cli.Fail(a...)
}
