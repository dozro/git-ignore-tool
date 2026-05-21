package commonStrings

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	GitTag    string
	GitCommit string
	BuildDate string
)

func fallBack() {
	if strings.TrimSpace(GitTag) != "" || strings.TrimSpace(GitCommit) != "" {
		return
	}
	info, _ := debug.ReadBuildInfo()
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			GitCommit = setting.Value
		case "vcs.time":
			BuildDate = setting.Value
		}
	}
	GitTag = "v0.0.0"
}

func VersionString() string {
	fallBack()
	return fmt.Sprintf("%s (%s, %s, default-build)", GitTag, GitCommit, BuildDate)
}
