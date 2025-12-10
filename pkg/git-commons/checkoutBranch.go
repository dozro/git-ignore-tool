package git_commons

import "os/exec"

func CheckoutBranch(gitExec, branchname string) error {
	cmd := exec.Command(gitExec, "checkout", branchname)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
