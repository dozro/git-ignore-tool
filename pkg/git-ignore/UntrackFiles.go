package git_ignore

import (
	"fmt"
	"os/exec"

	"git.gay/zakynthos/go-git-ignore/internal/pkg/commonStrings"
	gitcommons "git.gay/zakynthos/go-git-ignore/pkg/git-commons"
)

func UntrackFiles(toRemove []string, gitExecPath string, nocommit, commitEachOnItsown, onSepBranch bool, coauthors gitcommons.CoAuthors) error {
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
			Untrack(file, gitExecPath, nocommit, coauthors)
			continue
		}
		Untrack(file, gitExecPath, false, coauthors)
	}
	if !nocommit {
		git := gitcommons.Git{
			GitExec: &gitExecPath,
		}
		comMsg := "untracking ignored files"
		gitCommit := gitcommons.GitCommit{
			Git:       &git,
			Files:     []string{commonStrings.GitignoreFileName},
			Message:   &comMsg,
			CoAuthors: coauthors,
		}
		err := gitCommit.Commit()
		if err != nil {
			return fmt.Errorf("error committing files: %v", err)
		}
	}
	return nil
}

func Untrack(toRemove, gitExecPath string, nocommit bool, coauthors gitcommons.CoAuthors) error {
	cmd := exec.Command(gitExecPath, "rm", "--cached", toRemove)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error removing %v: %v, %v", toRemove, err, cmd.Err)
	}
	if !nocommit {
		git := gitcommons.Git{
			GitExec: &gitExecPath,
		}
		comMsg := "untracking ignored files"
		gitCommit := gitcommons.GitCommit{
			Git:       &git,
			Files:     []string{commonStrings.GitignoreFileName},
			Message:   &comMsg,
			CoAuthors: coauthors,
		}
		err := gitCommit.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
