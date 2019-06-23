package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

// Repo contains information about a git repository.
type Repo struct {
	os.FileInfo
}

// Repos returns all git repositories present in a directory.
func Repos(dir string) ([]Repo, error) {
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		_, err := os.Stat(path.Join(dir, item.Name(), ".git"))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		repos = append(repos, Repo{item})
	}

	return repos, nil
}

// Command runs a git command inside a repository.
func (r Repo) Command(commands []string, dir string) ([]byte, error) {
	cmd := exec.Command("git", commands...)
	cmd.Dir = path.Join(dir, r.Name())

	return cmd.CombinedOutput()
}
