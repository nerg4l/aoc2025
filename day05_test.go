package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay05(t *testing.T) {
	tests := []struct {
		name  string
		input io.ReadSeeker
		want1 string
		want2 string
	}{
		{
			name: "small",
			input: strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32
`),
			want1: `3
`,
			want2: `14
`,
		},
		{
			name:  "big",
			input: mayOpen("./day05.input"),
			want1: `623
`,
			want2: `353507173555373
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day05{}))
	}
}
