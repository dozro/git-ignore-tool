package main

import (
	"git-ignore/internal/pkg/commonStrings"
	git_commons "git-ignore/pkg/git-commons"
	gitignore "git-ignore/pkg/git-ignore"
	"os"

	log "github.com/sirupsen/logrus"
)

func ignore(workdir, gitExecPath string, nocreate, nocommit *bool, ignore []string) {
	var file *os.File
	if gitignore.FileExists(commonStrings.GitignoreFileName) {
		var err error
		file, err = os.OpenFile(commonStrings.GitignoreFileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Error reading gitignore: %v", err)
		}
	} else {
		log.Infof("Gitignore file %s not found, creating it", commonStrings.GitignoreFileName)
		if *nocreate {
			log.Fatalf("Will not create a gitignore if none exists")
		}
		file = gitignore.CreateNewGitIgnore(commonStrings.GitignoreFileName)
	}
	err := gitignore.AddToGitIgnore(file, ignore)
	if err != nil {
		log.Fatalf("Error adding gitignore: %v", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalf("Error closing gitignore: %v", err)
	}
	if !*nocommit {
		err = git_commons.AddToTracking(gitExecPath, commonStrings.GitignoreFileName)
		if err != nil {
			log.Fatalf("Error adding gitignore: %v", err)
		}
		err = git_commons.Commit(gitExecPath, commonStrings.GitignoreFileName, buildGitCommitMsgAfterIgnore(ignore))
		if err != nil {
			log.Fatalf("Error committing gitignore: %v", err)
		}
	}
	log.Infof("Added gitignore to %s", commonStrings.GitignoreFileName)
}

func buildGitCommitMsgAfterIgnore(ignores []string) string {
	return buildGitCommitMsg("ignoring files, using the git-ignore tool", "ignored files", ignores)
}
