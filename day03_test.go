package main

import (
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day03{},
			input: strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111
`),
			want1: `357
`,
			want2: `3121910778619
`,
		},
		{
			name:  "big",
			day:   Day03{},
			input: mayOpen("./day03.input"),
			want1: `17031
`,
			want2: `168575096286051
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
