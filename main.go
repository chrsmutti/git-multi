package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	WorkingDir string `short:"w" long:"working-dir" default:"." description:"set the working directory"`
	NoGroup    bool   `long:"no-group" description:"do not group same outputs"`
	Depth      int    `short:"d" long:"depth" default:"1" description:"depth of folders to look into for git repositories"`
	NoColor    bool   `long:"no-color" description:"do not print color characters"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)

	commands, err := parser.Parse()
	if err != nil {
		if !flags.WroteHelp(err) {
			parser.WriteHelp(os.Stderr)
		}

		os.Exit(1)
	}

	opts.WorkingDir = path.Clean(opts.WorkingDir)

	repos, err := Repos(opts.WorkingDir, 1, opts.Depth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get git repos: %v\n", err)
		os.Exit(1)
	}

	if err = run(commands, repos, opts.WorkingDir, !opts.NoGroup, !opts.NoColor); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run git %s: %v\n", strings.Join(commands, " "), err)
		os.Exit(1)
	}
}
