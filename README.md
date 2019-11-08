# git-multi

[![Actions Status](https://github.com/chrsmutti/git-multi/workflows/Go/badge.svg)](https://github.com/chrsmutti/git-multi/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A rewrite of [git-multi](https://github.com/tkrajina/git-plus/blob/master/git-multi) in Go.

Execute a single git command on multiple git repositories

## Installation

#### Standalone

`git-multi` can be easily installed as an executable. Download the latest
[compiled binaries](https://github.com/chrsmutti/git-multi/releases) and put it
anywhere in your executable path.

#### Source

Prerequisites for building from source are:

- `make`
- [Go](https://golang.org/doc/install)

Clone this repository and run make install:

```sh
git clone https://github.com/chrsmutti/git-multi

cd git-multi
make install
```

## Usage

```
Usage:
  git-multi [OPTIONS]

Application Options:
  -w, --working-dir= set the working directory (default: .)
      --no-group     do not group same outputs
  -d, --depth=       depth of folders to look into for git repositories
                     (default: 1)
      --no-color     do not print color characters

Help Options:
  -h, --help         Show this help message
```

If you have repositories ~/Projects/Repo1, ~/Projects/Repo2, ...:

```bash
cd ~/Projects

# Check the status of all repositories:
git multi status

# Which is the same as:
git multi

# You can also pass flags to git commands:
git multi status -- --short
```

The basic usage is simple:
`git multi <normal_git_commands_here> -- <normal_git_flags_here>`

## Group By Output

By default this version of git multi always groups by output, if that's not what
you desire you can use the `--no-group` flag.

```
git multi --no-group
```

# License

`git-multi` is licensed under the [MIT License](https://opensource.org/licenses/MIT).
