package git_commons

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func GetPathToGitExecutable() (string, error) {
	path, err := exec.LookPath("git") // Replace "git" with target executable name
	if err != nil {
		log.Error("Git not found in PATH")
		return "", err
	}
	return path, nil
}
