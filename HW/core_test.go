package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqAllLines_Success(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		opts  Options
		want  Result
	}{
		{
			name:  "basic unique keeps first appearance order",
			lines: []string{"a", "b", "a", "c", "b"},
			opts:  Options{},
			want: Result{
				Lines:  []string{"a", "b", "c"},
				Counts: []int{2, 2, 1},
			},
		},
		{
			name:  "-i ignore case",
			lines: []string{"A", "a", "Aa", "AA", "b"},
			opts:  Options{IgnoreCase: true},
			want: Result{
				Lines:  []string{"A", "Aa", "b"},
				Counts: []int{2, 2, 1},
			},
		},
		{
			name:  "-d only duplicates",
			lines: []string{"a", "b", "a", "c"},
			opts:  Options{OnlyDup: true},
			want: Result{
				Lines:  []string{"a"},
				Counts: []int{2},
			},
		},
		{
			name:  "-u only unique",
			lines: []string{"a", "b", "a", "c"},
			opts:  Options{OnlyUniq: true},
			want: Result{
				Lines:  []string{"b", "c"},
				Counts: []int{1, 1},
			},
		},
		{
			name: "-f 1 should group by tail after first field and not lose Thanks",
			lines: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			opts: Options{SkipFields: 1},
			want: Result{
				Lines:  []string{"We love music.", "", "I love music of Kartik.", "Thanks."},
				Counts: []int{3, 1, 2, 1},
			},
		},
		{
			name:  "-s 2 skip chars (UTF-8 safe in buildKey)",
			lines: []string{"abXYZ", "ab123", "abXYZ"},
			opts:  Options{SkipChars: 2},
			want: Result{
				Lines:  []string{"abXYZ", "ab123"},
				Counts: []int{2, 1},
			},
		},
		{
			name: "-i groups Kartik/kartik together",
			lines: []string{
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
			opts: Options{IgnoreCase: true},
			want: Result{
				Lines:  []string{"I love MuSIC of Kartik.", "Thanks."},
				Counts: []int{4, 1},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := UniqAllLines(tc.lines, tc.opts)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestUniqAllLines_Failures(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		opts  Options
	}{
		{
			name:  "negative SkipFields",
			lines: []string{"a"},
			opts:  Options{SkipFields: -1},
		},
		{
			name:  "negative SkipChars",
			lines: []string{"a"},
			opts:  Options{SkipChars: -5},
		},
		{
			name:  "conflicting flags -d and -u",
			lines: []string{"a", "a"},
			opts:  Options{OnlyDup: true, OnlyUniq: true},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := UniqAllLines(tc.lines, tc.opts)
			require.Error(t, err)
		})
	}
}
