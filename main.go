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

	repos, err := Repos(opts.WorkingDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get git repos: %v\n", err)
		os.Exit(1)
	}

	if err = run(commands, repos, opts.WorkingDir, !opts.NoGroup); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run git %s: %v\n", strings.Join(commands, " "), err)
		os.Exit(1)
	}
}
