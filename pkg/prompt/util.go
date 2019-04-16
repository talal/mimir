package prompt

import (
	"os"
	"strconv"

	"github.com/talal/go-bits/color"
)

func getBoolEnv(key string) bool {
	val, err := strconv.ParseBool(os.Getenv("MIMIR_KUBE"))
	if err != nil {
		return false
	}

	return val
}

func handleError(err error) {
	if err == nil {
		return
	}

	color.Fprintf(os.Stderr, color.Red, "Prompt error: %v\n", err)
}
