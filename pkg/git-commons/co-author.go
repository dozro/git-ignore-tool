package git_commons

import (
	"fmt"
	"strings"
)

type CoAuthor struct {
	Name  string
	Email string
	IsBot bool
}

func (a CoAuthor) CoAuthorString() string {
	if !a.IsBot {
		return fmt.Sprintf("%s <%s>", a.Name, a.Email)
	}
	return fmt.Sprintf("Co-authored-by: %s 🤖 <%s>", a.Name, a.Email)
}

type CoAuthors []CoAuthor

func (a CoAuthors) CoAuthorsString() string {
	var authors []string

	for _, author := range a {
		authors = append(authors, author.CoAuthorString())
	}

	return strings.Join(authors, "\n")
}
