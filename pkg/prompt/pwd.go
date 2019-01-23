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
	if err != nil {
		cwd = os.Getenv("PWD")
	}
	cwd = filepath.Clean(cwd)

	if cwd == "/" {
		return color.Sprintf(color.Blue, cwd)
	}

	nearestAccessiblePath := findNearestAccessiblePath(cwd)

	if nearestAccessiblePath != cwd {
		inAccessiblePath := strings.TrimPrefix(cwd, nearestAccessiblePath)
		nearestAccessiblePath = shortenLongPath(stripHomeDir(nearestAccessiblePath), 1)

		return color.Sprintf(color.Blue, nearestAccessiblePath) +
			color.Sprintf(color.Red, inAccessiblePath)
	}

	pathToDisplay := stripHomeDir(cwd)
	pathToDisplay = shortenLongPath(pathToDisplay, 2)

	gitDir, err := findGitRepo(cwd)
	handleError(err)

	if gitDir != "" {
		return color.Sprintf(color.Blue, pathToDisplay) + " " +
			color.Sprintf(color.Cyan, currentGitBranch(gitDir))
	}

	return color.Sprintf(color.Blue, pathToDisplay)
}

func findNearestAccessiblePath(path string) string {
	_, err := os.Stat(path)
	if err == nil {
		return path
	}

	return findNearestAccessiblePath(filepath.Dir(path))
}

func shortenLongPath(path string, length int) string {
	pList := strings.Split(path, "/")
	if len(pList) < 7 {
		return path
	}

	shortenedPList := pList[:len(pList)-length]
	for i, v := range shortenedPList {
		// shortenedPList[0] will be an empty string due to leading '/'
		if len(v) > 0 {
			shortenedPList[i] = v[:1]
		}
	}

	shortenedPList = append(shortenedPList, pList[len(pList)-length:]...)
	return strings.Join(shortenedPList, "/")
}

func stripHomeDir(path string) string {
	return strings.Replace(path, os.Getenv("HOME"), "~", 1)
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
