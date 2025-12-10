package main

import (
	"git-ignore/pkg/git-commons"
	"os"
	"regexp"
	"strings"

	sc "codeberg.org/dozrye/golang_simpleconfig"
	log "github.com/sirupsen/logrus"
)

var gitIgnoreExecNameRegexp = regexp.MustCompile("git-ignore$")
var gitUnIgnoreExecNameRegexp = regexp.MustCompile("git-unignore$")
var gitUntrackIgnoredRegexp = regexp.MustCompile("git-untrack-ignored$")

func main() {
	workdir, err := os.Getwd()
	gitExec, gitExecLookupErr := git_commons.GetPathToGitExecutable()
	if gitExecLookupErr != nil {
		log.Fatal(gitExecLookupErr)
	}
	calledAs := os.Args[0]
	sch := sc.SimpleConfigHandler{}
	sch.Init("GitIgnore", true, true, true, nil)
	nocreate := sch.GetBooleanOption(sc.ConfigEntry{Key: "NOCREATE", Description: "don't create a gitignore if none exists", DefaultBool: false})
	nocommit := sch.GetBooleanOption(sc.ConfigEntry{Key: "NOCOMMIT", Description: "don't commit a gitignore", DefaultBool: false})
	sch.ParseFlags()
	args := sch.GetCMDArgs(25)
	if err != nil {
		log.Fatal(err)
	}
	if !git_commons.IsGitRepo(workdir) {
		log.Fatalf("%s is not a git repo", workdir)
	}
	if gitIgnoreExecNameRegexp.MatchString(calledAs) {
		ignore(workdir, gitExec, nocreate, nocommit, args)
	} else if gitUnIgnoreExecNameRegexp.MatchString(calledAs) {
		unignore(workdir, gitExec, nocreate, nocommit, args)
	} else if gitUntrackIgnoredRegexp.MatchString(calledAs) {
		untrackignored(gitExec, nocommit)
	} else {
		log.Fatalf("this executable shouldn't be called with %s", calledAs)
	}
}

func buildGitCommitMsg(whatshort, intro string, ignores []string) string {
	var builder strings.Builder

	builder.WriteString(whatshort)
	builder.WriteString("\n\n")
	builder.WriteString(intro)
	builder.WriteString(":\n")
	builder.WriteString(strings.Join(ignores, "\n"))

	return builder.String()
}
