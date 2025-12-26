package main

import (
	"strings"
	"testing"
)

func TestDay11(t *testing.T) {
	tests := []dayTest{
		{
			name: "small-1",
			day:  Day11{},
			input: strings.NewReader(`aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`),
			want1: `5
`,
			want2: `0
`,
		},
		{
			name: "small-2",
			day:  Day11{},
			input: strings.NewReader(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`),
			want1: `0
`,
			want2: `2
`,
		},
		{
			name:  "big",
			day:   Day11{},
			input: mayOpen("./day11.input"),
			want1: `599
`,
			want2: `393474305030400
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
