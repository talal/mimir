package prompt

import (
	"os"

	"github.com/talal/go-bits/color"
)

func handleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.Red, "Prompt error: %v\n", err)
	}
}
