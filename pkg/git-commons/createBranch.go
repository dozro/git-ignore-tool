package git_commons

import "os/exec"

func CreateBranch(gitExec, branchname string) error {
	cmd := exec.Command(gitExec, "branch", branchname)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
