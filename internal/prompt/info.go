package prompt

import (
	"errors"
	"strings"
)

// Info returns the prompt info line.
func Info() string {
	cwd := getCwd()
	if cwd == "" {
		handleError(errors.New("could not get path for current working directory"))
	}

	info := make([]string, 0, 3)
	info = appendUnlessEmpty(info, getDir(cwd))
	info = appendUnlessEmpty(info, getKube())
	info = appendUnlessEmpty(info, getOSCloud())
	return strings.Join(info, " ")
}
