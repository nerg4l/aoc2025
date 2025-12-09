package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Day06 struct{}

func scanTable(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r == '\n' {
			return start + width, data[start : start+width], nil
		}
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) {
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func (Day06) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(scanTable)

	var rows [][]int
	var columns []int
	operationsI := 0

	result := 0

	for s.Scan() {
		l := s.Text()
		switch l {
		case "+":
			n := 0
			for _, row := range rows {
				n += row[operationsI]
			}
			result += n
			operationsI++
		case "*":
			n := 1
			for _, row := range rows {
				n *= row[operationsI]
			}
			result += n
			operationsI++
		case "\n":
			if operationsI > 0 {
				continue
			}
			rows = append(rows, columns)
			columns = make([]int, 0, len(columns))
			continue
		default:
			i, err := strconv.Atoi(l)
			if err != nil {
				return err
			}
			columns = append(columns, i)
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func (Day06) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	var rows []string

	for s.Scan() {
		rows = append(rows, s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	if len(rows) < 1 {
		return fmt.Errorf("invalid number of rows")
	}

	result := 0
	l := len(rows)
	numbers := make([]int, 0, l-1)
	var sb strings.Builder
	for i := len(rows[0]) - 1; i >= 0; i-- {
		for _, row := range rows[:l-1] {
			sb.WriteByte(row[i])
		}

		rawNum := strings.Trim(sb.String(), " ")
		if len(rawNum) == 0 {
			numbers = numbers[:0]
			continue
		}

		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return err
		}
		numbers = append(numbers, num)
		switch rows[l-1][i] {
		case '*':
			n := 1
			for _, num := range numbers {
				n *= num
			}
			result += n
			numbers = numbers[:0]
		case '+':
			n := 0
			for _, num := range numbers {
				n += num
			}
			result += n
			numbers = numbers[:0]
		}
		sb.Reset()
	}

	fmt.Fprintln(w, result)

	return nil
}
