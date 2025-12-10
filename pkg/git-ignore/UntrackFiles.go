package git_ignore

import (
	"fmt"
	gitcommons "git-ignore/pkg/git-commons"
	"os/exec"
)

func UntrackFiles(toRemove []string, gitExecPath string, nocommit, commitEachOnItsown, onSepBranch bool) error {
	if onSepBranch {
		if err := gitcommons.CreateBranch(gitExecPath, "untracking"); err != nil {
			return fmt.Errorf("error creating branch: %v", err)
		}
		if err := gitcommons.CheckoutBranch(gitExecPath, "untracking"); err != nil {
			return fmt.Errorf("error checking out branch: %v", err)
		}
	}
	for _, file := range toRemove {
		if commitEachOnItsown {
			Untrack(file, gitExecPath, nocommit)
			continue
		}
		Untrack(file, gitExecPath, false)
	}
	if !nocommit {
		err := gitcommons.CommitFiles(gitExecPath, toRemove, "untracking ignored files")
		if err != nil {
			return fmt.Errorf("error committing files: %v", err)
		}
	}
	return nil
}

func Untrack(toRemove, gitExecPath string, nocommit bool) error {
	cmd := exec.Command(gitExecPath, "rm", "--cached", toRemove)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error removing %v: %v, %v", toRemove, err, cmd.Err)
	}
	if !nocommit {
		err := gitcommons.Commit(gitExecPath, toRemove, "untracking ignored files")
		if err != nil {
			return err
		}
	}
	return nil
}
