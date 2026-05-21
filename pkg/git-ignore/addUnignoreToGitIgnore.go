package git_ignore

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func AddUnignoreToGitIgnore(gitignore *os.File, excludePatterns []string) error {
	existingIgnoreFile, err := ReadGitIgnore(gitignore)
	if err != nil {
		return err
	}

	// Build set for fast lookup
	existing := make(map[string]struct{}, len(existingIgnoreFile))
	for _, line := range existingIgnoreFile {
		existing[line] = struct{}{}
	}

	// Ensure we append
	_, err = gitignore.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	for _, pattern := range excludePatterns {
		if _, ok := existing[pattern]; ok {
			log.Errorf("Is explicitly ignored, pls remove the explicit ignore first: %s", pattern)
			continue
		}

		if _, ok := existing["!"+pattern]; ok {
			log.Infof("Is already unignored: %s", pattern)
			continue
		}

		log.Infof("Adding unignore pattern: %s", pattern)

		_, err := gitignore.WriteString("!" + pattern + "\n")
		if err != nil {
			return err
		}

		// update in-memory set so duplicates in same run are avoided
		existing[pattern] = struct{}{}
	}

	return nil
}
