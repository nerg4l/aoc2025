package main

import (
	"strings"
	"testing"
)

func TestDay07(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day07{},
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
			day:   Day07{},
			input: mayOpen("./day07.input"),
			want1: `1678
`,
			want2: `357525737893560
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
