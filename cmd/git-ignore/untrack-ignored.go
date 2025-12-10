package main

import (
	"git-ignore/internal/pkg/commonStrings"
	gitignore "git-ignore/pkg/git-ignore"
	"os"

	log "github.com/sirupsen/logrus"
)

func untrackignored(gitExecPath string, nocommit *bool) {
	var file *os.File
	if gitignore.FileExists(commonStrings.GitignoreFileName) {
		var err error
		file, err = os.OpenFile(commonStrings.GitignoreFileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Error reading gitignore: %v", err)
		}
	} else {
		log.Fatalf("gitignore file %s does not exist", commonStrings.GitignoreFileName)
	}
	defer file.Close()
	ignores, err := gitignore.ReadGitIgnore(file)
	if err != nil {
		log.Fatalf("Error reading gitignore: %v", err)
	}
	if err := gitignore.UntrackFiles(ignores, gitExecPath, *nocommit, *nocommit, true); err != nil {
		log.Fatalf("Error untracking files: %v", err)
	}
}
