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
	// Path is the relative path to the repo from the working dir.
	Path string
}

// Repos returns all git repositories present in a directory.
func Repos(dir string, current int, depth int) ([]Repo, error) {
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		pathName := path.Join(dir, item.Name())
		_, err := os.Stat(path.Join(pathName, ".git"))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		repos = append(repos, Repo{item, pathName})
	}

	if current < depth {
		for _, repo := range repos {
			nested, err := Repos(repo.Path, current+1, depth)
			if err != nil {
				return nil, err
			}

			repos = append(repos, nested...)
		}
	}

	return repos, nil
}

// Command runs a git command inside a repository.
func (r Repo) Command(commands []string, dir string, color bool) ([]byte, error) {
	if color {
		commands = append([]string{"-c", "color.ui=always"}, commands...)
	}

	cmd := exec.Command("git", commands...)
	cmd.Dir = path.Join(dir, r.Path)

	return cmd.CombinedOutput()
}
