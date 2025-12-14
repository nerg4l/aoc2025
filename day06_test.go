package main

import (
	"strings"
	"testing"
)

func TestDay06(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day06{},
			input: strings.NewReader(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`),
			want1: `4277556
`,
			want2: `3263827
`,
		},
		{
			name:  "big",
			day:   Day06{},
			input: mayOpen("./day06.input"),
			want1: `
`,
			want2: `
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
