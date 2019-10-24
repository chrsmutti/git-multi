package main

import (
	"fmt"
	"strings"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

type result struct {
	repo *Repo
	out  string
	err  error
}

func run(commands []string, repos []Repo, dir string, group bool, color bool) error {
	if len(commands) == 0 {
		commands = []string{"status"}
	}

	w, err := terminal.Width()
	if err != nil {
		return err
	}

	fmt.Println(strings.Repeat("-", int(w)))
	fmt.Printf("Executing git %s\n", strings.Join(commands, " "))
	fmt.Println(strings.Repeat("-", int(w)))

	results := make(chan result, len(repos))

	for _, repo := range repos {
		go func(r Repo) {
			out, err := r.Command(commands, dir, color)
			results <- result{&r, string(out), err}
		}(repo)
	}

	outputs, err := outputs(repos, results, group)
	if err != nil {
		return err
	}

	for _, output := range outputs {
		fmt.Printf("%s:\n", strings.Join(output.repos, ", "))

		lines := strings.Split(output.out, "\n")
		for _, line := range lines {
			fmt.Printf("\t%s\n", line)
		}
	}

	return nil
}
