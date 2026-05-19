package world

import "github.com/unmango/go/world/os"

type Os interface {
	os.Cmd
	os.Env
	os.Fs
	os.Id
	os.Net
	os.Stdio
	os.Sys
}

func SystemOs() Os {
	return os.System
}
