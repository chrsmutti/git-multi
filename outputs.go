package main

import (
	"os/exec"
)

type output struct {
	repos []string
	out   string
}

func outputs(repos []Repo, results chan result, group bool) ([]output, error) {
	seen := make(map[string][]result)

	for range repos {
		result := <-results

		if result.err != nil {
			if _, ok := result.err.(*exec.ExitError); !ok {
				return nil, result.err
			}
		}

		var key string
		if group {
			key = result.out
		} else {
			key = result.repo.Path
		}

		seen[key] = append(seen[key], result)
	}

	outputs := make([]output, 0)

	// Every result has the same out, if group is false, results should be a
	// slice with len = 1.
	for _, results := range seen {
		repos := make([]string, len(results))
		for i, result := range results {
			repos[i] = result.repo.Path
		}

		outputs = append(outputs, output{repos, results[0].out})
	}

	return outputs, nil
}
