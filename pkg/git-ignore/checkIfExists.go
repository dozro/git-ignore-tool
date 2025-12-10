package git_ignore

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		log.Debugf("Error checking if file exists: %v", err)
		return false
	}
}
