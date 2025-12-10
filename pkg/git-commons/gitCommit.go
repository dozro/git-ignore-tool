package git_commons

import (
	"os/exec"
	"strings"
)

func AddToTracking(gitExec, toCommit string) error {
	cmd := exec.Command(gitExec, "add", toCommit)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Commit(gitExec, toCommit, commitMsg string) error {
	cmd := exec.Command(gitExec, "commit", "-m", commitMsg, toCommit)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func CommitFiles(gitExec string, toCommit []string, commitMsg string) error {
	cmd := exec.Command(gitExec, "commit", "-m", commitMsg, strings.Join(toCommit, " "))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
