package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

type Day09 struct{}

type Rectangle struct {
	A, B Point
}

func (r Rectangle) Area() int {
	return (max(r.A.X, r.B.X) - min(r.A.X, r.B.X) + 1) *
		(max(r.A.Y, r.B.Y) - min(r.A.Y, r.B.Y) + 1)
}

func (r Rectangle) MinX() int {
	return min(r.A.X, r.B.X)
}

func (r Rectangle) MaxX() int {
	return max(r.A.X, r.B.X)
}

func (r Rectangle) MinY() int {
	return min(r.A.Y, r.B.Y)
}

func (r Rectangle) MaxY() int {
	return max(r.A.Y, r.B.Y)
}

func (Day09) Points(s *bufio.Scanner) iter.Seq2[Point, error] {
	return func(yield func(Point, error) bool) {
		for s.Scan() {
			l := s.Text()
			str := strings.SplitN(l, ",", 2)
			if len(str) != 2 {
				yield(Point{}, fmt.Errorf("missing `-` from input %q", l))
				return
			}
			coord := Point{}
			coord.X, _ = strconv.Atoi(str[0])
			coord.Y, _ = strconv.Atoi(str[1])
			if !yield(coord, nil) {
				return
			}
		}
	}
}

func (d Day09) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	var coordinates []Point

	largest := -1

	for a, err := range d.Points(s) {
		if err != nil {
			return err
		}

		for _, b := range coordinates {
			r := Rectangle{A: a, B: b}
			area := r.Area()
			if largest < area {
				largest = area
			}
		}

		coordinates = append(coordinates, a)
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w, largest)

	return nil
}

func (d Day09) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	var coordinates []Point

	for coord, err := range d.Points(s) {
		if err != nil {
			return err
		}
		coordinates = append(coordinates, coord)
	}

	if err := s.Err(); err != nil {
		return err
	}
	largest := -1

	for i, a := range coordinates[:len(coordinates)-1] {
	loop:
		for _, b := range coordinates[i+1:] {
			r := Rectangle{A: a, B: b}

			perim := ((r.MaxX() - r.MinX()) + (r.MaxY() - r.MinY())) * 2
			walkedPerim := 0

			for j, k := 0, len(coordinates)-1; j < len(coordinates); j++ {
				start, end := coordinates[k], coordinates[j]
				k = j

				if start.X < end.X { // top
					if end.X <= r.MinX() || r.MaxX() <= start.X {
						continue
					}
					y := start.Y
					if y <= r.MinY() {
						walkedPerim += min(end.X, r.MaxX()) - max(start.X, r.MinX())
					} else if y >= r.MaxY() {
						walkedPerim -= min(end.X, r.MaxX()) - max(start.X, r.MinX())
					} else {
						continue loop
					}
				} else if start.Y < end.Y { // right
					if end.Y <= r.MinY() || r.MaxY() <= start.Y {
						continue
					}
					x := start.X
					if x >= r.MaxX() {
						walkedPerim += min(end.Y, r.MaxY()) - max(start.Y, r.MinY())
					} else if x <= r.MinX() {
						walkedPerim -= min(end.Y, r.MaxY()) - max(start.Y, r.MinY())
					} else {
						continue loop
					}
				} else if end.X < start.X { // bottom
					if start.X <= r.MinX() || r.MaxX() <= end.X {
						continue
					}
					y := start.Y
					if y >= r.MaxY() {
						walkedPerim += min(start.X, r.MaxX()) - max(end.X, r.MinX())
					} else if y <= r.MinY() {
						walkedPerim -= min(start.X, r.MaxX()) - max(end.X, r.MinX())
					} else {
						continue loop
					}
				} else if end.Y < start.Y { // left
					if start.Y <= r.MinY() || r.MaxY() <= end.Y {
						continue
					}
					x := start.X
					if x <= r.MinX() {
						walkedPerim += min(start.Y, r.MaxY()) - max(end.Y, r.MinY())
					} else if x >= r.MaxX() {
						walkedPerim -= min(start.Y, r.MaxY()) - max(end.Y, r.MinY())
					} else {
						continue loop
					}
				}
			}

			if walkedPerim != perim {
				continue
			}

			area := r.Area()
			if largest < area {
				largest = area
			}
		}
	}

	fmt.Fprintln(w, largest)

	return nil
}
