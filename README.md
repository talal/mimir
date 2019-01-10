# Mímir

[![GitHub release](https://img.shields.io/github/release/talal/mimir.svg)](https://github.com/talal/mimir/releases/latest)
[![Build Status](https://travis-ci.org/talal/mimir.svg?branch=master)](https://travis-ci.org/talal/mimir)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/mimir)](https://goreportcard.com/report/github.com/talal/mimir)

Mímir is a fast and minimal Zsh prompt in [Go](https://golang.org). The look of
Mímir is inspired by [Pure](https://github.com/sindresorhus/pure) and the
functionality is inspired by
[prettyprompt](https://github.com/majewsky/gofu#prettyprompt).

![screenshot](./screenshot.png)

Features:
- It is very fast: loads in ~5ms with everything turned on.
- Long directory paths are shortened (see screenshot above).
- Kubernetes context and namespace info is shown using the list of kubernetes
  configuration file(s) taken from `KUBECONFIG` environment variable.
- OpenStack cloud info is shown using the `CURRENT_OS_CLOUD` environment
  variable, if specified, otherwise the standard OpenStack environment
  variables are used to show the cloud info.

The Mímir Go binary only displays the (pre) prompt line with all the
information. The actual prompt line (with the symbol; `❯` as shown in the
screenshot above) is configured in the shell's config file. This allows for
flexibility:
- You can configure the prompt to your preference, e.g. custom prompt symbol,
  user or host name before prompt symbol, etc.
- You can use this with any shell of your choosing. The description says Bash
  and Zsh because these are the shells which I have tested Mímir on, but in
  reality you can use Mímir with any shell as long as it allows you to load a
  binary as a prompt.

## Installation

### Pre-compiled binaries

Pre-compiled binaries for Linux and macOS are avaiable on the
[releases](https://github.com/talal/mimir/releases/latest) page.

The binaries are static executables.

### Homebrew

```
brew install talal/tap/mimir
```

### Building from source

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/mimir
$ cd $GOPATH/src/github.com/talal/mimir
$ make install
```

this will put the binary in `/usr/bin/mimir` or `/usr/local/bin/mimir` for macOS.

## Usage

The following usage examples for Bash and Zsh are just one example of how Mímir
can be configured. The examples below will result in a setup similar to the
screenshot shown above: the prompt symbol (`❯`) changes to red if the previous
command exited with an error.

### Bash

Add this to your `.bashrc` file:

```bash
prompt_mimir_cmd() {
  if [ $? != 0 ]; then
    local prompt_symbol="\[\e[0;31m\]❯\[\e[0m\]"
  else
    local prompt_symbol="\[\e[0;35m\]❯\[\e[0m\]"
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

### Options

| Option | Description | Usage |
| --- | --- | --- |
| `MIMIR_KUBE` | Disable Kubernetes context and namespace info. | `export MIMIR_KUBE='false'` |
| `MIMIR_OS_CLOUD` | Disable OpenStack cloud info. | `export MIMIR_OS_CLOUD='false'` |

## Credits

Most of the source code is borrowed from
[prettyprompt](https://github.com/majewsky/gofu#prettyprompt).
