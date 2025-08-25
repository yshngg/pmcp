package version

import (
	"strings"
	"testing"
)

func TestInfoString(t *testing.T) {
	testCases := []struct {
		name     string
		info     info
		expected []string
	}{
		{
			name: "all fields set",
			info: info{
				Number:    "1.2.3",
				GitCommit: "abcdef1",
				BuildDate: "2025-08-21T12:00:00Z",
			},
			expected: []string{
				"Version:  v1.2.3",
				"Commit:   abcdef1",
				"Build:    2025-08-21T12:00:00Z",
			},
		},
		{
			name: "missing fields",
			info: info{
				Number:    "",
				GitCommit: "",
				BuildDate: "",
			},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.info.String()
			for _, sub := range tc.expected {
				if !strings.Contains(out, sub) {
					t.Errorf("expected output to contain %q, got %q", sub, out)
				}
			}
		})
	}
}

func TestInfoSet(t *testing.T) {
	var i info

	i.Set("v2.0.0", "deadbeef", "2025-08-21T15:00:00Z")
	if i.Number != "2.0.0" {
		t.Errorf("expected Number to be '2.0.0', got %q", i.Number)
	}
	if i.GitCommit != "deadbeef" {
		t.Errorf("expected GitCommit to be 'deadbeef', got %q", i.GitCommit)
	}
	if i.BuildDate != "2025-08-21T15:00:00Z" {
		t.Errorf("expected BuildDate to be '2025-08-21T15:00:00Z', got %q", i.BuildDate)
	}

	// Test default values
	i.Set("", "", "")
	if i.Number != "(unknown)" {
		t.Errorf("expected Number to be '(unknown)', got %q", i.Number)
	}
	if i.GitCommit != "" {
		t.Errorf("expected GitCommit to be '', got %q", i.GitCommit)
	}
	if len(i.BuildDate) == 0 {
		t.Error("expected BuildDate to be set to current time, got empty string")
	}
}
