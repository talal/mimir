package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		os.Stderr.Write([]byte("\x1B[1;31mPrompt error: " + "no arguing with Mímir" + "\x1B[0m\n"))
		os.Exit(1)
	}

	var line []string
	line = appendUnlessEmpty(line, getDir())
	line = appendUnlessEmpty(line, getKube())
	line = appendUnlessEmpty(line, getOSCloud())

	os.Stdout.Write([]byte("\n" + strings.Join(line, " ") + "\n"))
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
