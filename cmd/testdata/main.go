package main

import (
	"os"

	"github.com/unmango/go/cmd"
)

func main() {
	var a []any
	for _, x := range os.Args[1:] {
		a = append(a, x)
	}

	cmd.Fail(a...)
}
