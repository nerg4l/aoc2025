package main

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			input: strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`),
			want1: `3
`,
			want2: `6
`,
		},
		{
			name:  "big",
			input: mayOpen("./day01.input"),
			want1: `982
`,
			want2: `6106
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day01{}))
	}
}
