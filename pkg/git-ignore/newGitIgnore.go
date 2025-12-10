package git_ignore

import (
	"git-ignore/internal/pkg/commonStrings"
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateNewGitIgnore(gitignoreFileName string) *os.File {
	f, err := os.Create(gitignoreFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(commonStrings.GitignoreNewGitIgnoreHeading)
	return f
}
