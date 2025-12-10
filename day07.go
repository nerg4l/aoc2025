package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Day07 struct{}

func (Day07) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	beams := map[int]struct{}{}

	for s.Scan() {
		l := s.Text()
		i := strings.Index(l, "S")
		if i != -1 {
			beams[i] = struct{}{}
			break
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	splits := 0
	var splitted []int
	for s.Scan() {
		l := s.Text()
		for i := range beams {
			if l[i] == '^' {
				splitted = append(splitted, i)
			}
		}
		for _, i := range splitted {
			delete(beams, i)
			beams[i-1] = struct{}{}
			beams[i+1] = struct{}{}
		}
		splits += len(splitted)
		splitted = splitted[:0]
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, splits)

	return nil
}

func (Day07) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	beams := map[int]int{}

	for s.Scan() {
		l := s.Text()
		i := strings.Index(l, "S")
		if i != -1 {
			beams[i] = 1
			break
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	var splitted []int
	for s.Scan() {
		l := s.Text()
		for i := range beams {
			if l[i] == '^' {
				splitted = append(splitted, i)
			}
		}
		for _, i := range splitted {
			splits := beams[i]
			delete(beams, i)
			beams[i-1] += splits
			beams[i+1] += splits
		}
		splitted = splitted[:0]
	}

	if err := s.Err(); err != nil {
		return err
	}

	splits := 0
	for _, i := range beams {
		splits += i
	}

	fmt.Fprintln(w, splits)

	return nil
}
