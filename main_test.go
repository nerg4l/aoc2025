package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type dayTest struct {
	name  string
	day   Day
	input io.ReadSeeker
	want1 string
	want2 string
}

func dayTester(tt dayTest, d Day) func(t *testing.T) {
	return func(t *testing.T) {
		if tt.input == nil {
			t.Skip("Input was not provided")
		}

		var w bytes.Buffer
		err := d.Segment1(tt.input, &w)
		if err != nil {
			t.Errorf("Segemnt1() got error %s", err)
		}
		if got := w.String(); got != tt.want1 {
			t.Errorf("Segemnt1() got %q, want %q", got, tt.want1)
		}

		tt.input.Seek(0, io.SeekStart)
		w.Reset()
		if err := d.Segment2(tt.input, &w); err != nil {
			t.Errorf("Segemnt2() got error %s", err)
		}
		if got := w.String(); got != tt.want2 {
			t.Errorf("Segemnt2() got %q, want %q", got, tt.want2)
		}
	}
}

func mayOpen(name string) io.ReadSeeker {
	finput, err := os.Open(name)
	if err != nil {
		return nil
	}
	return finput
}
