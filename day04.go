package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
)

type Day04 struct{}

type Point struct {
	X, Y int
}

func (p Point) Neighbours() []Point {
	return []Point{
		{X: p.X, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y - 1},
	}
}

func (Day04) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	_, candidates := movableRolls(s)

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, len(candidates))

	return nil
}

func (Day04) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	adjacentRolls, candidates := movableRolls(s)

	allCandidates := map[Point]struct{}{}
	maps.Copy(allCandidates, candidates)

	for {
		found := false
		next := map[Point]struct{}{}

		for p := range candidates {
			for _, np := range p.Neighbours() {
				if _, ok := adjacentRolls[np]; !ok {
					continue
				}
				if _, ok := allCandidates[np]; ok {
					continue
				}
				adjacentRolls[np] -= 1
				if adjacentRolls[np] < 4 {
					found = true
					next[np] = struct{}{}
				}
			}
		}
		candidates = next
		maps.Copy(allCandidates, candidates)
		for p := range candidates {
			delete(adjacentRolls, p)
		}

		if !found {
			break
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, len(allCandidates))

	return nil
}

func movableRolls(s *bufio.Scanner) (adjacentRolls map[Point]int, candidates map[Point]struct{}) {
	tmpAdjacentRolls := map[Point]int{}
	adjacentRolls = map[Point]int{}
	candidates = map[Point]struct{}{}

	for y := 0; s.Scan(); y++ {
		l := s.Text()
		for x, r := range l {
			p := Point{X: x, Y: y}

			if r != '@' {
				continue
			}
			adjacentRolls[p] = 0

			if tmpAdjacentRolls[p] < 4 {
				candidates[p] = struct{}{}
			}

			for _, np := range p.Neighbours() {
				tmpAdjacentRolls[np] += 1
				if tmpAdjacentRolls[np] > 3 {
					delete(candidates, np)
				}
			}
		}
	}

	for p := range adjacentRolls {
		adjacentRolls[p] = tmpAdjacentRolls[p]
	}

	return adjacentRolls, candidates
}
