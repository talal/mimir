# Mímir

[![GitHub Release](https://img.shields.io/github/release/talal/mimir.svg?style=flat-square)](https://github.com/talal/mimir/releases/latest)
[![Build Status](https://img.shields.io/travis/talal/mimir/master.svg?style=flat-square)](https://travis-ci.org/talal/mimir)
[![Go Report Card](https://goreportcard.com/badge/github.com/talal/mimir?style=flat-square)](https://goreportcard.com/report/github.com/talal/mimir)
[![Software License](https://img.shields.io/github/license/talal/mimir.svg?style=flat-square)](LICENSE)

Mímir is a fast and minimal shell prompt written in [Go](https://golang.org).
The look of Mímir is inspired by [Pure](https://github.com/sindresorhus/pure)
and the functionality is inspired by
[prettyprompt](https://github.com/majewsky/gofu#prettyprompt).

[![asciicast](https://asciinema.org/a/225675.svg)](https://asciinema.org/a/225675)

The Mímir Go binary only displays the (pre) prompt line with all the
information. The actual prompt line used for input is configured in the shell's
config file. This allows for flexibility:
- You can configure the prompt line to your preference, e.g. you can specify a
  prompt symbol of your choice, add user or host name info, etc.
- You can use Mímir with any shell of your choosing. The description says Bash
  and Zsh because these are the shells which I have tested Mímir on but
  technically you can use Mímir with any shell as long as it allows you to load
  a binary as a prompt.

## Features

- It is very fast: loads under 5ms with everything turned on
  ([benchmark](https://asciinema.org/a/225680)).
- Shows current Git branch name.
- Long directory paths are shortened and inaccessible paths are highlighted in
  red.
- Kubernetes context and namespace info is shown using `KUBECONFIG` environment
  variable. If multiple config files are specified in the variable value then
  the first one with current context info is used.
- OpenStack cloud info is shown using the standard OpenStack environment
  variables that begin with `OS_`. Variables that specify IDs are used if name
  specifying variables are not available, e.g. `OS_PROJECT_DOMAIN_ID` would be
  used if `OS_PROJECT_DOMAIN_NAME` is not available. You can override the cloud
  info by exporting the `CURRENT_OS_CLOUD` variable with an arbitrary value.


## Installation

### Installer script

The simplest way to install Mímir on Linux or macOS is to run:

```
$ sh -c "$(curl -sL git.io/getmimir)"
```

This will put the binary in `/usr/local/bin/mimir`

### Pre-compiled binaries

Pre-compiled binaries for Linux and macOS are avaiable on the
[releases](https://github.com/talal/mimir/releases/latest) page.

The binaries are static executables.

### Homebrew

```
$ brew install talal/tap/mimir
```

### Building from source

The only required build dependency is [Go](https://golang.org/).

```
$ go get github.com/talal/mimir
$ cd $GOPATH/src/github.com/talal/mimir
$ make install
```

This will put the binary in `/usr/local/bin/mimir`

## Usage

The following usage examples are just one example of how Mímir can be
configured. The examples below will result in a setup similar to the one shown
in the demo above: the prompt symbol (`❯`) changes to red if the previous
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
| `MIMIR_KUBE` | Disable Kubernetes context and namespace info. | `export MIMIR_KUBE='0'` |
| `MIMIR_OS_CLOUD` | Disable OpenStack cloud info. | `export MIMIR_OS_CLOUD='0'` |

## Credits

Most of the source code is borrowed from
[prettyprompt](https://github.com/majewsky/gofu#prettyprompt).
