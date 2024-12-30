package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
)

func main() {
	v, err := version.Sprint(
		context.Background(),
		os.Args[0],
	)
	if err != nil {
		cmd.Fail(err)
	} else {
		fmt.Println(v)
	}
}
