package main

import (
	"strings"
	"testing"
)

func TestDay04(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day04{},
			input: strings.NewReader(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`),
			want1: `13
`,
			want2: `43
`,
		},
		{
			name:  "big",
			day:   Day04{},
			input: mayOpen("./day04.input"),
			want1: `1489
`,
			want2: `
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
