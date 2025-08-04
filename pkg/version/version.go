package version

import "fmt"

type info struct {
	Number    string
	GitCommit string
	BuildDate string
}

func (i info) String() string {
	return fmt.Sprintf("PMCP Version: v%s\nGit Commit: %s\nBuild Date: %s", i.Number, i.GitCommit, i.BuildDate)
}

var (
	Number    string
	GitCommit string
	BuildDate string

	Info = info{
		Number:    "0.1.0-dev",
		GitCommit: "HEAD",
		BuildDate: "2006-01-02T15:04:05Z07:00",
	}
)

func init() {
	if len(Number) != 0 {
		Info.Number = Number
		Info.GitCommit = GitCommit
		Info.BuildDate = BuildDate
	}
}
