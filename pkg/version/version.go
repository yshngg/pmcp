package version

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

const (
	NumberUnknown = "(unknown)"
	NumberDevel   = "(devel)"
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
		v := i.Number
		if v != NumberUnknown && v != NumberDevel {
			v = "v" + v
		}
		if _, err := builder.WriteString(fmt.Sprintf("Version:\t%s\n", v)); err != nil {
			panic(err)
		}
	}
	if len(i.GitCommit) != 0 {
		if _, err := builder.WriteString(fmt.Sprintf("Commit:\t%s\n", i.GitCommit)); err != nil {
			panic(err)
		}
	}
	if len(i.BuildDate) != 0 {
		if _, err := builder.WriteString(fmt.Sprintf("Build:\t%s", i.BuildDate)); err != nil {
			panic(err)
		}
	}
	return builder.String()
}

func (i *info) Set(versionNumber, gitCommit, buildDate string) {
	if len(versionNumber) == 0 {
		versionNumber = NumberUnknown
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
		if v := buildInfo.Main.Version; len(v) != 0 && v != NumberDevel {
			Number = v
		}

		for _, setting := range buildInfo.Settings {
			if setting.Key == "vcs.revision" {
				GitCommit = setting.Value
				if len(GitCommit) > 7 {
					GitCommit = GitCommit[:7]
				}
				break
			}
		}
	}
	Info.Set(Number, GitCommit, BuildDate)
}
