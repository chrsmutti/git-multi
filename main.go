package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var workingDir = flag.String("w", ".", "set the working directory")
var depth = flag.Int("d", 1, "depth of folders to look into for git repositories")
var noGroup = flag.Bool("no-group", false, "do not group same outputs")
var noColor = flag.Bool("no-color", false, "do not print color characters")

func main() {
	flag.Parse()
	var commands = flag.Args()
	var workingDir = path.Clean(*workingDir)

	repos, err := Repos(workingDir, 1, *depth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get git repos: %v\n", err)
		os.Exit(1)
	}

	if err = run(commands, repos, workingDir, !*noGroup, !*noColor); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run git %s: %v\n", strings.Join(commands, " "), err)
		os.Exit(1)
	}
}
