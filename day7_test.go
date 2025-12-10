package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	tests := []struct {
		name  string
		input io.ReadSeeker
		want1 string
		want2 string
	}{
		{
			name: "small",
			input: strings.NewReader(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`),
			want1: `21
`,
			want2: `40
`,
		},
		{
			name:  "big",
			input: mayOpen("./day7.input"),
			want1: `1678
`,
			want2: `357525737893560
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day7{}))
	}
}
