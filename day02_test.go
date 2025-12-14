package main

import (
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	tests := []dayTest{
		{
			name: "small",
			day:  Day02{},
			input: strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
`),
			want1: `1227775554
`,
			want2: `4174379265
`,
		},
		{
			name:  "big",
			day:   Day02{},
			input: mayOpen("./day02.input"),
			want1: `19386344315
`,
			want2: `34421651192
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, tt.day))
	}
}
