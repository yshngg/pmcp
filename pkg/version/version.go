package version

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

type info struct {
	Number    string
	GitCommit string
	BuildDate string
}

func (i info) String() string {
	builder := strings.Builder{}
	if _, err := builder.WriteString("Prometheus Model Context Protocol Server\n"); err != nil {
		panic(err)
	}
	if len(i.Number) != 0 {
		if _, err := builder.WriteString(fmt.Sprintf("Version: v%s\n", i.Number)); err != nil {
			panic(err)
		}
	}
	if len(i.GitCommit) != 0 {
		if _, err := builder.WriteString(fmt.Sprintf("Commit: %s\n", i.GitCommit)); err != nil {
			panic(err)
		}
	}
	if len(i.BuildDate) != 0 {
		if _, err := builder.WriteString(fmt.Sprintf("Build: %s", i.BuildDate)); err != nil {
			panic(err)
		}
	}
	return builder.String()
}

func (i *info) Set(versionNumber, gitCommit, buildDate string) {
	if len(versionNumber) == 0 {
		versionNumber = "(unknown)"
	}
	i.Number, _ = strings.CutPrefix(versionNumber, "v")

	i.GitCommit = gitCommit

	if len(buildDate) == 0 {
		buildDate = time.Now().UTC().Format(time.RFC3339)
	}
	i.BuildDate = buildDate
}

var (
	Number    string
	GitCommit string
	BuildDate string

	Info = info{
		Number:    "(unknown)",
		GitCommit: "",
		BuildDate: "",
	}
)

func init() {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if len(buildInfo.Main.Version) != 0 {
			Number = buildInfo.Main.Version
		}

		settings := buildInfo.Settings
		for _, setting := range settings {
			if setting.Key == "vcs.revision" {
				GitCommit = setting.Value[:7]
				break
			}
		}
	}
	Info.Set(Number, GitCommit, BuildDate)
}
