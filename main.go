package main

import (
	"os"

	"github.com/talal/go-bits/color"
	"github.com/talal/mimir/internal/prompt"
)

func main() {
	if len(os.Args) > 1 {
		color.Fprintln(os.Stderr, color.Red, "Prompt error: no arguing with Mímir")
		os.Exit(1)
	}

	os.Stdout.Write([]byte("\n" + prompt.Info() + "\n"))
}
