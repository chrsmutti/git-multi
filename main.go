package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	WorkingDir string `short:"w" long:"working-dir" description:"set the working directory"`
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

	if len(opts.WorkingDir) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get working dir: %v\n", err)
			os.Exit(1)
		}

		opts.WorkingDir = cwd
	}

	opts.WorkingDir = strings.TrimSuffix(strings.TrimSpace(opts.WorkingDir), "/")

	repos, err := gitRepos(opts.WorkingDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get git repos: %v\n", err)
		os.Exit(1)
	}

	if err = runGitCmd(commands, repos, opts.WorkingDir); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run git command %s: %v\n", "git "+strings.Join(commands, " "), err)
		os.Exit(1)
	}
}

func runGitCmd(commands []string, repos []os.FileInfo, workingDir string) error {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Executing %s\n", "git "+strings.Join(commands, " "))
	fmt.Println("--------------------------------------------------------------------------------")

	type result struct {
		repo string
		out  []byte
		err  error
	}

	results := make(chan result, len(repos))

	run := func(repo os.FileInfo) {
		cmd := exec.Command("git", commands...)
		cmd.Dir = workingDir + "/" + repo.Name()

		out, err := cmd.CombinedOutput()

		results <- result{repo.Name(), out, err}
	}

	for _, repo := range repos {
		go run(repo)
	}

	grouped := make(map[string][]string)

	for range repos {
		result := <-results

		if result.err != nil {
			if _, ok := result.err.(*exec.ExitError); !ok {
				return result.err
			}
		}

		grouped[string(result.out)] = append(grouped[string(result.out)], result.repo)
	}

	for output, repos := range grouped {
		fmt.Printf("%s:\n", strings.Join(repos, ", "))

		lines := strings.Split(output, "\n")
		for _, line := range lines {
			fmt.Printf("\t%s\n", line)
		}
	}

	return nil
}

func gitRepos(path string) ([]os.FileInfo, error) {
	items, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	n := 0
	repos := make([]os.FileInfo, len(items))

	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		_, err := os.Stat(path + "/" + item.Name() + "/.git")
		if err != nil && os.IsNotExist(err) {
			// not a git repo
			continue
		} else if err != nil {
			return nil, err
		}

		repos[n] = item
		n++
	}

	return repos[:n], nil
}
