package version

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"strings"
	"text/tabwriter"
	"time"
)

const (
	NumberUnknown = "(unknown)"
	NumberDevel   = "(devel)"

	GitCommitLength = 7
)

type number string

func (n number) String() string {
	if len(n) != 0 && n != NumberUnknown && n != NumberDevel {
		n = "v" + n
	}
	return string(n)
}

type info struct {
	Number    number
	GitCommit string
	BuildDate string
}

func (i info) String() string {
	var buff bytes.Buffer
	w := tabwriter.NewWriter(&buff, 0, 2, 2, ' ', 0)

	if len(i.Number) != 0 {
		if _, err := fmt.Fprintf(w, "Version:\t%s\n", i.Number); err != nil {
			panic(err)
		}
	}

	if len(i.GitCommit) != 0 {
		if _, err := fmt.Fprintf(w, "Commit:\t%s\n", i.GitCommit); err != nil {
			panic(err)
		}
	}
	if len(i.BuildDate) != 0 {
		if _, err := fmt.Fprintf(w, "Build:\t%s\n", i.BuildDate); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}
	return buff.String()
}

func (i *info) Set(versionNumber, gitCommit, buildDate string) {
	if len(versionNumber) == 0 {
		versionNumber = NumberUnknown
	}
	numberStr, _ := strings.CutPrefix(string(versionNumber), "v")
	i.Number = number(numberStr)

	i.GitCommit = gitCommit

	if len(buildDate) == 0 {
		buildDate = time.Now().UTC().Format(time.RFC3339)
	}
	i.BuildDate = buildDate
}

var (
	Number    string
	GitCommit string = "$Id$" // ref: https://git-scm.com/docs/gitattributes#_ident
	BuildDate string

	Info = info{
		Number:    NumberUnknown,
		GitCommit: "",
		BuildDate: "",
	}
)

func init() {
	defer func() { Info.Set(Number, GitCommit, BuildDate) }()

	// Parse ident expansion: "$Id: <oid> $"
	if strings.HasPrefix(GitCommit, "$Id: ") && strings.HasSuffix(GitCommit, " $") {
		GitCommit = GitCommit[5 : len(GitCommit)-2]
	}
	if GitCommit == "$Id"+"$" {
		GitCommit = ""
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if len(Number) == 0 {
			if v := buildInfo.Main.Version; len(v) != 0 {
				Number = v
			}
		}

		// Prefer the real commit SHA from build info over ident (blob OID).
		for _, setting := range buildInfo.Settings {
			if setting.Key == "vcs.revision" && len(setting.Value) != 0 {
				GitCommit = setting.Value
				break
			}
		}
	}

	// Normalize to short form for display.
	if len(GitCommit) > GitCommitLength {
		GitCommit = GitCommit[:GitCommitLength]
	}
}
