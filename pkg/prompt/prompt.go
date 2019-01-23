package prompt

import (
	"os"
	"strings"
)

// Info returns the prompt info line.
func Info() string {
	var line []string
	line = appendUnlessEmpty(line, getDir())
	if os.Getenv("MIMIR_KUBE") != "0" {
		line = appendUnlessEmpty(line, getKube())
	}
	if os.Getenv("MIMIR_OS_CLOUD") != "0" {
		line = appendUnlessEmpty(line, getOSCloud())
	}

	return strings.Join(line, " ")
}

func appendUnlessEmpty(list []string, val string) []string {
	if val == "" {
		return list
	}

	return append(list, val)
}
