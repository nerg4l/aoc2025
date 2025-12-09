package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct{}

type IDRange struct {
	start, end int
}

func (Day05) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	ranges := []IDRange{}

	var err error
	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			break
		}
		r := IDRange{}
		i := strings.Index(l, "-")
		if i == -1 {
			return fmt.Errorf("missing `-` from input %q", l)
		}
		r.start, err = strconv.Atoi(l[:i])
		if err != nil {
			return err
		}
		r.end, err = strconv.Atoi(l[i+1:])
		if err != nil {
			return err
		}

		ranges = append(ranges, r)
	}

	if err := s.Err(); err != nil {
		return err
	}

	result := 0

	for s.Scan() {
		id, err := strconv.Atoi(s.Text())
		if err != nil {
			return err
		}
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				result++
				break
			}
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func (Day05) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	ranges := []IDRange{}

	var err error
	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			break
		}
		r := IDRange{}
		i := strings.Index(l, "-")
		if i == -1 {
			return fmt.Errorf("missing `-` from input %q", l)
		}
		r.start, err = strconv.Atoi(l[:i])
		if err != nil {
			return err
		}
		r.end, err = strconv.Atoi(l[i+1:])
		if err != nil {
			return err
		}

		ranges = append(ranges, r)
	}
	slices.SortFunc(ranges, func(a, b IDRange) int {
		return a.start - b.start
	})

	n := 0
	result := 0
	for _, r := range ranges {
		if n > r.end {
			continue
		}
		if n < r.start {
			n = r.start
		}
		result += r.end - n + 1
		n = r.end + 1
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}
