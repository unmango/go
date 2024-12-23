package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/go/devops/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "devops",
	Short: "Helper utilities for developing code",
}

func main() {
	log.SetLevel(log.ErrorLevel)

	rootCmd.AddCommand(
		cmd.NewList(&cmd.ListOptions{}),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
