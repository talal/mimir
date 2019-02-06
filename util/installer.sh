#!/bin/sh
# script borrowed from https://github.com/getantibody/installer
set -e
DOWNLOAD_URL="https://github.com/talal/mimir/releases/download"

last_version() {
  curl -s https://raw.githubusercontent.com/talal/homebrew-tap/master/Formula/mimir.rb |
    grep version |
    cut -f2 -d'"'
}

download() {
	version="$(last_version)" || true
  test -z "$version" && {
    echo "Unable to get mimir version."
    exit 1
  }
  echo "Downloading mimir v$version for $(uname -s)..."
  rm -f /tmp/mimir /tmp/mimir.tar.gz
  curl -s -L -o /tmp/mimir.tar.gz \
    "$DOWNLOAD_URL/v$version/mimir-$version-$(uname -s)_amd64.tar.gz"
}

extract() {
  tar -xf /tmp/mimir.tar.gz -C /tmp
}

main() {
	download
	extract || true
	sudo mv -f /tmp/mimir /usr/local/bin/mimir
	rm -f /tmp/mimir.tar.gz
	echo "mimir v$version installed in $(which mimir)"
}

main
