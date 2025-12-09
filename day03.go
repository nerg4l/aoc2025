package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Day03 struct{}

func (Day03) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	result := 0

	for s.Scan() {
		l := s.Text()
		var (
			major  rune
			majorI = -1
			minor  rune
		)
		for i, r := range l {
			if r > major {
				major = r
				majorI = i
			}
		}
		var remainder string
		if majorI == len(l)-1 {
			remainder = l[:len(l)-1]
		} else {
			remainder = l[majorI+1:]
		}
		minor = rune(remainder[0])
		for _, r := range remainder {
			if r > minor {
				minor = r
			}
		}
		if majorI == len(l)-1 {
			major, minor = minor, major
		}
		v, err := strconv.Atoi(string(major) + string(minor))
		if err != nil {
			return fmt.Errorf("invalid character found:%q, %s", l, err)
		}
		result += v
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func (Day03) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	result := 0

	for s.Scan() {
		l := s.Text()
		v, err := strconv.Atoi(highestJoltage(l, 12))
		if err != nil {
			return fmt.Errorf("invalid character found:%q, %s", l, err)
		}
		result += v
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func highestJoltage(bank string, batteries int) string {
	if len(bank) == 1 {
		return bank
	}

	var (
		major  rune
		majorI = -1
	)
	for i, r := range bank {
		if r > major {
			major = r
			majorI = i
		}
	}

	result := string(major)

	if batteries <= 1 {
		return result
	}

	if majorI < len(bank)-1 {
		result += highestJoltage(bank[majorI+1:], batteries-1)
	}

	if l := len(result); l < batteries && len(bank) > l {
		return highestJoltage(bank[:len(bank)-l], batteries-l) + result
	}

	return result
}
