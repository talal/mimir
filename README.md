# Mímir

[![GitHub release](https://img.shields.io/github/release/talal/mimir.svg)](https://github.com/talal/mimir/releases/latest)
[![Build Status](https://travis-ci.org/talal/mimir.svg?branch=master)](https://travis-ci.org/talal/mimir)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/mimir)](https://goreportcard.com/report/github.com/talal/mimir)

Mímir is a fast and minimal Zsh prompt in [Go](https://golang.org). The look of Mímir is inspired by [Pure](https://github.com/sindresorhus/pure) and the functionality is inspired by [prettyprompt](https://github.com/majewsky/gofu#prettyprompt).

![screenshot](./screenshot.png)

Features:
- Long directory paths are shortened (see screenshot above).
- Kubernetes context/namespace info is shown using the list of kubernetes configuration file(s) taken from `KUBECONFIG` environment variable.
- OpenStack cloud info is shown using the `CURRENT_OS_CLOUD` environment variable.

## Installation

Download the latest pre-compiled binary from the [releases](https://github.com/talal/mimir/releases/latest) page.

Alternatively, you can also build from source:

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/mimir
$ cd $GOPATH/src/github.com/talal/mimir
$ make install
```

this will put the binary in `/usr/bin/mimir` or `/usr/local/bin/mimir` for macOS.

## Usage

### Bash

Add this to your `.bashrc` file:

```bash
prompt_mimir_cmd() {
  local exit_code="$?"
  local resetColor='\[\e[0m\]'
  local red='\[\e[0;31m\]'
  local magenta='\[\e[0;35m\]'

  if [ $exit_code != 0 ]; then
    local prompt_symbol="${red}❯${resetColor}"
  else
    local prompt_symbol="${magenta}❯${resetColor}"
  fi

  PS1="$(/path/to/mimir)\n${prompt_symbol} "
}
PROMPT_COMMAND=prompt_mimir_cmd
```

### Zsh

Add this to your `.zshrc` file:

```zsh
autoload -Uz add-zsh-hook
prompt_mimir_cmd() { /path/to/mimir }
add-zsh-hook precmd prompt_mimir_cmd

prompt_symbol='❯'
PROMPT='%(?.%F{magenta}.%F{red})${prompt_symbol}%f '
```

## Credits

Most of the source code is borrowed from [prettyprompt](https://github.com/majewsky/gofu#prettyprompt).
