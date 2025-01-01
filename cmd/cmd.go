package cmd

import (
	"fmt"
	"os"
)

func Fail(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
