package main

import (
	"os"
	"os/exec"
)

func main() {
	tag, _ := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	commit, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	date, _ := exec.Command("date").Output()

	os.WriteFile(".generated-version-tag.txt", tag, 0644)
	os.WriteFile(".generated-version-commit.txt", commit, 0644)
	os.WriteFile(".generated-version-build-date.txt", date, 0644)
}
