package git_commons

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func IsGitRepo(basePath string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/.git", basePath)); err == nil {
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		log.Debugf("Error checking if in git repo: %v", err)
		return false
	}
}
