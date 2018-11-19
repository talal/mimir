package prompt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

// GetDir returns the current working directory and the
// git branch (if the directory is a git repo).
func GetDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		handleError(err)
	}

	homePath := "/Users/" + os.Getenv("USER")
	displayPath := strings.Replace(cwd, homePath, "~", 1)

	if pathList := strings.Split(displayPath, "/"); len(pathList) > 6 {
		for i, v := range pathList[:len(pathList)-2] {
			pathList[i] = v[:1]
		}
		displayPath = strings.Join(pathList, "/")
	}

	gitDir, err := findRepo(cwd)
	if err != nil {
		handleError(err)
	}

	if gitDir != "" {
		return color.Sprintf(color.Blue, displayPath) + " " +
			color.Sprintf(color.Cyan, getRepo(gitDir))
	}

	return color.Sprintf(color.Blue, displayPath)
}

func findRepo(path string) (string, error) {
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
		return findRepo(filepath.Dir(path))
	}

	if !fi.IsDir() {
		return "", nil
	}

	return gitEntry, nil
}

func getRepo(gitDir string) string {
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
