package init

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	util "github.com/unmango/go/cmd"
	"github.com/unmango/go/devops/version"
	"github.com/unmango/go/devops/work"
)

var (
	AutoVersionSource   = "auto"
	GitHubVersionSource = "github"

	VersionSources = []string{
		AutoVersionSource,
		GitHubVersionSource,
	}
)

type VersionOptions struct {
	work.ChdirOptions
	Name   string
	Source string
}

type VersionArgs []string

func (v VersionArgs) Source() string {
	return v[0]
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use:   "version [dependency]",
		Short: "Generates files for versioning the specified dependency",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				util.Fail(err)
			}

			var (
				vargs = VersionArgs(args)
				src   version.Source
				err   error
			)

			switch opts.Source {
			case AutoVersionSource:
				src, err = version.GuessSource(vargs.Source())
			case GitHubVersionSource:
				src = version.GitHub(vargs.Source())
			default:
				err = fmt.Errorf("unsupported source: %s", opts.Source)
			}
			if err != nil {
				util.Fail(err)
			}

			if err = version.Init(ctx, opts.Name, src); err != nil {
				util.Fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	cmd.Flags().StringVarP(&opts.Source, "source", "s", AutoVersionSource,
		fmt.Sprintf("source of dependency, one of: [%s]", strings.Join(VersionSources, ", ")),
	)
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "explicit version name")

	return cmd
}
