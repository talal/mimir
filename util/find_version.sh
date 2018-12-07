#!/bin/sh
awk '$1 == "##" { if (/Unreleased/) { print "dev" } else { print $2 } }' CHANGELOG.md | sed 's/^v//' | head -n1
