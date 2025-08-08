package utils

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	now := time.Now().Truncate(time.Second)

	testCases := []struct {
		name     string
		timeStr  string
		expected time.Time
		err      error
	}{
		{
			"RFC3339",
			"2006-01-02T15:04:05Z",
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			nil,
		},
		{
			"Unix Timestamp",
			strconv.FormatInt(now.Unix(), 10),
			now,
			nil,
		},
		{
			"Empty String",
			"",
			time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			InvalidFormatError(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := ParseTime(tc.timeStr)
			require.Equal(t, tc.err, err)
			require.Truef(t, tc.expected.Equal(actual), "expected: %s, actual: %s", tc.expected, actual)
		})
	}
}
