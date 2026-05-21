package git_commons

import (
	"fmt"
	"os/exec"
)

type Git struct {
	GitExec *string
}

func (a *Git) createCommand(commandArgs ...string) (*exec.Cmd, error) {
	if a.GitExec == nil {
		return nil, fmt.Errorf("no Git exec specified")
	}
	return exec.Command(*a.GitExec, commandArgs...), nil
}
