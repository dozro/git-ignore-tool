package git_ignore

import (
	"fmt"
	"git-ignore/internal/pkg/commonStrings"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func AddToGitIgnore(gitignore *os.File, excludePatterns []string) error {
	existingIgnoreFile, err := ReadGitIgnore(gitignore)
	if err != nil {
		return err
	}

	for _, pattern := range excludePatterns {
		var existsAlready bool
		patRegexp, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}
		for _, existingIgnore := range existingIgnoreFile {
			if patRegexp.MatchString(existingIgnore) {
				existsAlready = true
				break
			}
		}
		if existsAlready {
			log.Infof("Skipping ignore pattern, as it already is ignored: %s", pattern)
			continue
		}
		log.Infof("Adding ignore pattern: %s", pattern)
		ignoreStr := fmt.Sprintf("\n%s\n%s\n", commonStrings.GitignoreComment, pattern)
		_, err = gitignore.WriteString(ignoreStr)
		if err != nil {
			return err
		}
	}

	return nil
}
