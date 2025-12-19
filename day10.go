package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

type Day10 struct{}

type Manual struct {
	TargetLights     int
	ButtonWiringBits []int
	ButtonWiring     [][]int
	TargetJoltage    []int
	raw              string
}

func (Day10) Manuals(s *bufio.Scanner) iter.Seq2[Manual, error] {
	return func(yield func(Manual, error) bool) {
		for s.Scan() {
			l := s.Text()
			strs := strings.Split(l, " ")
			m := Manual{raw: l}
			if len(strs) < 3 {
				yield(m, fmt.Errorf("invalid manual %q", l))
				return
			}

			for _, str := range strs {
				switch str[0] {
				case '[':
					for i, r := range str[1:(len(str) - 1)] {
						if r == '#' {
							m.TargetLights |= 1 << i
						}
					}
				case '(':
					var wiringBits int
					var wiring []int
					for _, button := range strings.Split(str[1:(len(str)-1)], ",") {
						x, err := strconv.Atoi(button)
						if err != nil {
							yield(m, err)
							return
						}
						wiringBits |= 1 << x
						wiring = append(wiring, x)
					}
					m.ButtonWiringBits = append(m.ButtonWiringBits, wiringBits)
					m.ButtonWiring = append(m.ButtonWiring, wiring)
				case '{':
					for _, joltage := range strings.Split(str[1:(len(str)-1)], ",") {
						x, err := strconv.Atoi(joltage)
						if err != nil {
							yield(m, err)
							return
						}
						m.TargetJoltage = append(m.TargetJoltage, x)
					}
				}
			}

			if !yield(m, nil) {
				return
			}
		}
	}
}

func (d Day10) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	result := 0

	for m, err := range d.Manuals(s) {
		if err != nil {
			return err
		}
		result += findFewestButtonPressForTargetLight(m)
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func findFewestButtonPressForTargetLight(m Manual) int {
	next := make([]int, len(m.ButtonWiringBits))
	for i, w := range m.ButtonWiringBits {
		if m.TargetLights == w {
			return 1
		}
		next[i] = w
	}
	for i := 2; ; i++ {
		buttons := next
		next = make([]int, 0, (len(buttons) * len(m.ButtonWiringBits)))
		for _, w := range buttons {
			for _, ww := range m.ButtonWiringBits {
				state := w ^ ww
				if m.TargetLights == state {
					return i
				}
				next = append(next, state)
			}
		}
	}
}

func (d Day10) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	result := 0

	for m, err := range d.Manuals(s) {
		if err != nil {
			return err
		}
		result += findFewestButtonPressForTagetJoltage(m)
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func findFewestButtonPressForTagetJoltage(m Manual) int {
	return -1
}
