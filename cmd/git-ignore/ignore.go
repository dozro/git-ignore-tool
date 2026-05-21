package main

import (
	"fmt"
	"os"

	gitcommons "git.gay/zakynthos/go-git-ignore/pkg/git-commons"
	gitignore "git.gay/zakynthos/go-git-ignore/pkg/git-ignore"

	"git.gay/zakynthos/go-git-ignore/internal/pkg/commonStrings"

	log "github.com/sirupsen/logrus"
)

func ignore(args IgnoreArgs) {
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
	err := gitignore.AddToGitIgnore(file, args.IgnorePatterns)
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
		comMsg := buildGitCommitMsgAfterIgnore(args.IgnorePatterns)
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
		if err = gitignore.UntrackFiles(args.IgnorePatterns, *args.GitExec, *args.NoCommit, *args.NoCommit, false, []gitcommons.CoAuthor{coAuthor}); err != nil {
			log.Fatalf("Error untracking files: %v", err)
		}
	}
	log.Infof("Added gitignore to %s", commonStrings.GitignoreFileName)
}

func buildGitCommitMsgAfterIgnore(ignores []string) string {
	return buildGitCommitMsg("ignoring files, using the git-ignore tool", "ignored files", ignores)
}
