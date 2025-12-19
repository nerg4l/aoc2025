package main

import (
	"strings"
	"testing"
)

func TestDay10(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day10{},
			input: strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`),
			want1: `7
`,
			want2: `33
`,
		},
		{
			name:  "big",
			day:   Day10{},
			input: mayOpen("./day10.input"),
			want1: `578
`,
			want2: `
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
