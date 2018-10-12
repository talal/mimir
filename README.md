# Mímir

[![Build Status](https://travis-ci.org/talal/mimir.svg?branch=master)](https://travis-ci.org/talal/mimir) [![Go Report Card](https://goreportcard.com/badge/github.com/talal/mimir)](https://goreportcard.com/report/github.com/talal/mimir)

Mímir is a fast and minimal Bash/Zsh prompt in [Go](https://golang.org). The look of Mímir is inspired by [Pure](https://github.com/sindresorhus/pure) and the functionality is inspired by [Gofu](gofu).

![screenshot](./screenshot.png)

Features:
- Long directory paths are shortened (see screenshot above).
- Kubernetes context/namespace info, context info is taken from `KUBECONFIG` environment variable and namespace info is taken from a `~/.kubectl-namespace` text file. I use [this shell function](https://gist.github.com/talal/7182c4d7c8f544fd590fffea903a0dae) to change the namespace and update it in the text file.
- OpenStack cloud info is shown using the `CURRENT_OS_CLOUD` environment variable.

## Installation/Usage

Prerequisites:
- `go`

```
$ go get github.com/talal/mimir
$ cd $GOPATH/src/github.com/talal/mimir
$ make install
```

Add this to your .bashrc/.zshrc file:

```
PROMPT="\$(/usr/local/bin/mimir \$?)"
```

## Credits

Most of the source code is borrowed from [Gofu](gofu).

[gofu]: https://github.com/majewsky/gofu
