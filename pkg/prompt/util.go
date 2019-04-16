package prompt

import (
	"os"

	"github.com/talal/go-bits/color"
)

func handleError(err error) {
	if err == nil {
		return
	}

	color.Fprintf(os.Stderr, color.Red, "Prompt error: %v\n", err)
}
