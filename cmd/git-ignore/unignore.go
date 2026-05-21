package main

import (
	"fmt"
	"os"

	"git.gay/zakynthos/go-git-ignore/internal/pkg/commonStrings"
	gitcommons "git.gay/zakynthos/go-git-ignore/pkg/git-commons"
	gitignore "git.gay/zakynthos/go-git-ignore/pkg/git-ignore"
	log "github.com/sirupsen/logrus"
)

func unignore(args IgnoreArgs) {
	var file *os.File
	if gitignore.FileExists(commonStrings.GitignoreFileName) {
		var err error
		file, err = os.OpenFile(commonStrings.GitignoreFileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Error reading gitignore: %v", err)
		}
	} else {
		log.Infof("Gitignore file %s not found, creating it", commonStrings.GitignoreFileName)
		if *args.NoCreate {
			log.Fatalf("Will not create a gitignore if none exists")
		}
		file = gitignore.CreateNewGitIgnore(commonStrings.GitignoreFileName)
	}
	err := gitignore.AddUnignoreToGitIgnore(file, args.IgnorePatterns)
	if err != nil {
		log.Fatalf("Error adding gitignore: %v", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalf("Error closing gitignore: %v", err)
	}
	git := gitcommons.Git{
		GitExec: args.GitExec,
	}
	var coAuthor gitcommons.CoAuthor
	if !*args.NoCoAuthor {
		coAuthor = gitcommons.CoAuthor{
			Name:  fmt.Sprintf("Git Ignore (%s)", commonStrings.VersionString()),
			Email: "git-ignore@cisnt.fyi",
			IsBot: true,
		}
	}

	if !*args.NoCommit {
		err = gitcommons.AddToTracking(*args.GitExec, commonStrings.GitignoreFileName)
		if err != nil {
			log.Fatalf("Error adding gitignore: %v", err)
		}
		comMsg := buildGitCommitMsgAfterUnIgnore(args.IgnorePatterns)
		gitCommit := gitcommons.GitCommit{
			Git:       &git,
			Files:     []string{commonStrings.GitignoreFileName},
			Message:   &comMsg,
			CoAuthors: []gitcommons.CoAuthor{coAuthor},
		}
		err = gitCommit.Commit()
		if err != nil {
			log.Fatalf("Error committing gitignore: %v", err)
		}
	}
	log.Infof("Added unignore to gitignore to %s", commonStrings.GitignoreFileName)
}

func buildGitCommitMsgAfterUnIgnore(ignores []string) string {
	return buildGitCommitMsg("unignoring files (explicit exceptions), using the git-ignore tool", "new explicit exceptions", ignores)
}
