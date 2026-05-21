package commonStrings

import (
	_ "embed"
	"fmt"
)

//go:generate go run ./scripts/version.go
//go:embed .generated-version-tag.txt
var GitTag string

//go:embed .generated-version-commit.txt
var GitCommit string

//go:embed .generated-version-build-date.txt
var BuildDate string

func VersionString() string {
	return fmt.Sprintf("%s (%s, %s)", GitTag, GitCommit, BuildDate)
}
