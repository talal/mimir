package prompt

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/talal/go-bits/color"
)

func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = os.Getenv("PWD")
	}
	return filepath.Clean(cwd)
}

func getBoolEnv(key string) bool {
	val, err := strconv.ParseBool(os.Getenv("MIMIR_KUBE"))
	if err != nil {
		return false
	}

	return val
}

func appendUnlessEmpty(list []string, val string) []string {
	if val == "" {
		return list
	}

	return append(list, val)
}

func handleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.Red, "Prompt error: %v\n", err)
	}
}
