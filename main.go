package main

import (
	"os"
	"strings"
)

func main() {
	var line []string
	line = appendUnlessEmpty(line, getDir())
	line = appendUnlessEmpty(line, getKube())
	line = appendUnlessEmpty(line, getOSCloud())

	os.Stdout.Write([]byte("\n" + strings.Join(line, " ") + "\n"))
	os.Stdout.Write([]byte(withColor(magenta, "❯ ")))
}

func handleError(err error) {
	if err != nil {
		os.Stderr.Write([]byte("\x1B[1;31mPrompt error: " + err.Error() + "\x1B[0m\n"))
	}
}

func appendUnlessEmpty(list []string, val string) []string {
	if val == "" {
		return list
	}

	return append(list, val)
}
