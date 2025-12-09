package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day02 struct{}

func ScanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := bytes.IndexByte(data, ',')
	if i == -1 {
		if !atEOF {
			return 0, nil, nil
		}
		// If we have reached the end, return the last token.
		return 0, bytes.Trim(data, "\n"), bufio.ErrFinalToken
	}
	// Return the token before the comma.
	return i + 1, bytes.Trim(data[:i], "\n"), nil
}

func (Day02) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(ScanComma)

	result := 0

	for s.Scan() {
		l := s.Text()
		i := strings.IndexByte(l, '-')
		if i == -1 {
			return fmt.Errorf("missing `-` from input %q", l)
		}
		start, err := strconv.Atoi(l[:i])
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(l[i+1:])
		if err != nil {
			return err
		}
		for id := start; id <= end; id++ {
			str := strconv.Itoa(id)
			if len(str)%2 != 0 {
				continue
			}
			j := len(str) / 2
			if str[:j] != str[j:] {
				continue
			}
			result += id
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)
	return nil
}

func (Day02) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(ScanComma)

	result := 0

	for s.Scan() {
		l := s.Text()
		i := strings.IndexByte(l, '-')
		if i == -1 {
			return fmt.Errorf("missing `-` from input %q", l)
		}
		start, err := strconv.Atoi(l[:i])
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(l[i+1:])
		if err != nil {
			return err
		}
		for id := start; id <= end; id++ {
			if !invalidID(id) {
				continue
			}
			result += id
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)
	return nil
}

func invalidID(id int) bool {
	str := strconv.Itoa(id)
	n := len(str) / 2
loop:
	for i := n; i > 0; i-- {
		x := str[:i]
		if len(str)%i != 0 {
			continue
		}
		for j := 0; j < len(str); j += i {
			if x != str[j:j+i] {
				continue loop
			}
		}
		return true
	}
	return false
}
