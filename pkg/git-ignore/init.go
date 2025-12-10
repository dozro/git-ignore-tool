package git_ignore

import "regexp"

func init() {
	gitignoreComments = regexp.MustCompile("#.*$")
}
