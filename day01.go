package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Day01 struct{}

func (Day01) Segment1(r io.Reader, w io.Writer) error {
	dial := 50

	zeros := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		clicks, err := strconv.Atoi(l[1:])
		if err != nil {
			return err
		}
		if l[0] == 'L' {
			dial -= clicks
		} else {
			dial += clicks
		}
		if dial%100 == 0 {
			zeros++
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	fmt.Fprintln(w, zeros)
	return nil
}

func (Day01) Segment2(r io.Reader, w io.Writer) error {
	dial := 50

	zeros := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		clicks, err := strconv.Atoi(l[1:])
		if err != nil {
			return err
		}
		prev := dial
		if l[0] == 'L' {
			dial -= clicks
		} else {
			dial += clicks
		}
		// When dial is zero or the rotation caused a sign change
		if (dial == 0) || (prev > 0 && dial < 0) || (prev < 0 && dial > 0) {
			zeros++
		}
		if dial > 99 {
			zeros += dial / 100
			dial %= 100
		}
		if dial < -99 {
			zeros += -dial / 100
			dial %= 100
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	fmt.Fprintln(w, zeros)
	return nil
}
