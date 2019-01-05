package prompt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

// getDir returns the current working directory and the git branch (if the
// directory is a git repo).
func getDir() string {
	cwd, err := os.Getwd()
	handleError(err)

	pathToDisplay := filepath.Clean(cwd)
	if pathToDisplay != "/" {
		homePath := os.Getenv("HOME")
		pathToDisplay = strings.Replace(pathToDisplay, homePath, "~", 1)

		if pathList := strings.Split(pathToDisplay, "/"); len(pathList) > 6 {
			for i, v := range pathList[:len(pathList)-2] {
				// pathList[0] will be an empty string due to leading '/'
				if len(v) > 0 {
					pathList[i] = v[:1]
				}
			}
			pathToDisplay = strings.Join(pathList, "/")
		}
	}

	gitDir, err := findGitRepo(cwd)
	handleError(err)

	if gitDir != "" {
		return color.Sprintf(color.Blue, pathToDisplay) + " " +
			color.Sprintf(color.Cyan, currentGitBranch(gitDir))
	}

	return color.Sprintf(color.Blue, pathToDisplay)
}

func findGitRepo(path string) (string, error) {
	gitEntry := filepath.Join(path, ".git")
	fi, err := os.Stat(gitEntry)
	switch {
	case err == nil:
		//found - continue below with further checks
	case !os.IsNotExist(err):
		return "", err
	case path == "/":
		return "", nil
	default:
		return findGitRepo(filepath.Dir(path))
	}

	if !fi.IsDir() {
		return "", nil
	}

	return gitEntry, nil
}

func currentGitBranch(gitDir string) string {
	bytes, err := ioutil.ReadFile(filepath.Join(gitDir, "HEAD"))
	if err != nil {
		handleError(err)
		return "unknown"
	}
	refSpec := strings.TrimSpace(string(bytes))

	// detached HEAD?
	if !strings.HasPrefix(refSpec, "ref: refs/") {
		return "detached"
	}

	refSpecDisplay := strings.TrimPrefix(refSpec, "ref: refs/heads/")
	return refSpecDisplay
}
