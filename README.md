# git-multi

A rewrite of [git-multi](https://github.com/tkrajina/git-plus/blob/master/git-multi) in Go.

Execute a single git command on multiple git repositories

## Usage

```
Usage:
  git-multi [OPTIONS]

Application Options:
  -w, --working-dir= set the working directory (default: .)
      --no-group     do not group same outputs

Help Options:
  -h, --help         Show this help message
```

If you have repositories ~/Projects/Repo1, ~/Projects/Repo2, ...:

```bash
cd ~/projects

# Check the status of all repositories:
git multi status

# Which is the same as:
git multi

# You can also pass flags to git commands:
git multi status -- --short
```

The basic usage is simple:
`git multi normal_git_commands_here -- normal_git_flags_here`

## Installation

You can use `go get` to install this tool.

```
go get -u github.com/chrsmutti/git-multi
```

## Group By Output

By default this version of git multi always groups by output, if that's not what
you desire you can use the `--no-group` flag.

```
git multi --no-group
```