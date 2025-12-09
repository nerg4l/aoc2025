package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	tests := []struct {
		name  string
		input io.ReadSeeker
		want1 string
		want2 string
	}{
		{
			name: "small",
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
			input: mayOpen("./day03.input"),
			want1: `17031
`,
			want2: `168575096286051
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day03{}))
	}
}
