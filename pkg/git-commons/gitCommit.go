package git_commons

import (
	"fmt"
	"os/exec"
	"strings"
)

type GitAdd struct {
	Git   *Git
	Files []string
}

func (a *GitAdd) GitAdd(files ...string) error {
	a.Files = files
	return a.Add()
}

func (a *GitAdd) createCommand() (*exec.Cmd, error) {
	if a.Git == nil {
		return nil, fmt.Errorf("git not configured")
	}
	return a.Git.createCommand("add", strings.Join(a.Files, " "))
}

func (a *GitAdd) Add() error {
	if a.Files == nil {
		return fmt.Errorf("git add: no files")
	}
	cmd, err := a.createCommand()
	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func AddToTracking(gitExec, toCommit string) error {
	cmd := exec.Command(gitExec, "add", toCommit)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

type GitCommit struct {
	Git       *Git
	Files     []string
	Message   *string
	CoAuthors CoAuthors
}

func (c *GitCommit) createCommand() (*exec.Cmd, error) {
	comMsg := *c.Message
	if c.Git == nil {
		return nil, fmt.Errorf("git not configured")
	}
	if c.CoAuthors != nil {
		comMsg = fmt.Sprintf("%s\n\n%s", comMsg, c.CoAuthors.CoAuthorsString())
	}
	return c.Git.createCommand("commit", "-m", comMsg, strings.Join(c.Files, " "))
}

func (c *GitCommit) Commit() error {
	if c.Git == nil {
		return fmt.Errorf("git not configured")
	}
	if c.Message == nil {
		return fmt.Errorf("message not configured")
	}
	if c.Files == nil {
		return fmt.Errorf("files to commit not configured")
	}
	cmd, err := c.createCommand()
	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
