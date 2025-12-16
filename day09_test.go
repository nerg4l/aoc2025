package main

import (
	"strings"
	"testing"
)

func TestDay09(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day09{},
			input: strings.NewReader(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`),
			want1: `50
`,
			want2: `24
`,
		},
		{
			name:  "big",
			day:   Day09{},
			input: mayOpen("./day09.input"),
			want1: `4752484112
`,
			want2: `1465767840
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
