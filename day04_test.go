package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay04(t *testing.T) {
	tests := []struct {
		name  string
		input io.ReadSeeker
		want1 string
		want2 string
	}{
		{
			name: "small",
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
			input: mayOpen("./day04.input"),
			want1: `1489
`,
			want2: `
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day04{}))
	}
}
