package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day08 struct {
	Pairs int
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) Distance(other Point3D) float64 {
	return math.Sqrt(math.Pow(float64(p.X)-float64(other.X), 2) +
		math.Pow(float64(p.Y)-float64(other.Y), 2) +
		math.Pow(float64(p.Z)-float64(other.Z), 2))
}

type JunctionCable struct {
	A, B Point3D
}

func (c JunctionCable) Len() float64 {
	return math.Sqrt(math.Pow(float64(c.A.X)-float64(c.B.X), 2) +
		math.Pow(float64(c.A.Y)-float64(c.B.Y), 2) +
		math.Pow(float64(c.A.Z)-float64(c.B.Z), 2))
}

// MinHeap implements heap.Interface for a min heap of Nodes
type MinHeap []JunctionCable

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Len() < h[j].Len() }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(JunctionCable))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (d Day08) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	var boxes []Point3D

	circuits := map[Point3D][]Point3D{}
	roots := map[Point3D]Point3D{}

	minH := &MinHeap{}
	heap.Init(minH)
	for s.Scan() {
		strs := strings.SplitN(s.Text(), ",", 3)
		if len(strs) != 3 {
			return fmt.Errorf("invalid line: %s", strs)
		}

		box := Point3D{}
		box.Y, _ = strconv.Atoi(strs[1])
		box.X, _ = strconv.Atoi(strs[0])
		box.Z, _ = strconv.Atoi(strs[2])

		for _, b := range boxes {
			heap.Push(minH, JunctionCable{A: b, B: box})
		}
		circuits[box] = []Point3D{box}
		roots[box] = box

		boxes = append(boxes, box)
	}

	if err := s.Err(); err != nil {
		return err
	}

	for i := 0; i < d.Pairs && minH.Len() > 0; i++ {
		c := heap.Pop(minH).(JunctionCable)
		pA, pB := roots[c.A], roots[c.B]
		if pA == pB {
			continue
		}
		extra := circuits[pB]
		delete(circuits, pB)
		circuits[pA] = append(circuits[pA], extra...)
		for _, C := range extra {
			roots[C] = pA
		}
	}

	var counts []int
	for _, v := range circuits {
		counts = append(counts, len(v))
	}
	slices.SortFunc(counts, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	sum := 1
	for _, v := range counts[:min(3, len(counts))] {
		sum *= v
	}

	fmt.Fprintln(w, sum)

	return nil
}

func (Day08) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	var boxes []Point3D

	circuits := map[Point3D][]Point3D{}
	roots := map[Point3D]Point3D{}

	minH := &MinHeap{}
	heap.Init(minH)
	for s.Scan() {
		strs := strings.SplitN(s.Text(), ",", 3)
		if len(strs) != 3 {
			return fmt.Errorf("invalid line: %s", strs)
		}

		box := Point3D{}
		box.Y, _ = strconv.Atoi(strs[1])
		box.X, _ = strconv.Atoi(strs[0])
		box.Z, _ = strconv.Atoi(strs[2])

		for _, b := range boxes {
			heap.Push(minH, JunctionCable{A: b, B: box})
		}
		circuits[box] = []Point3D{box}
		roots[box] = box

		boxes = append(boxes, box)
	}

	if err := s.Err(); err != nil {
		return err
	}

	var found *JunctionCable
	for minH.Len() > 0 {
		c := heap.Pop(minH).(JunctionCable)
		pA, pB := roots[c.A], roots[c.B]
		if pA == pB {
			continue
		}
		extra := circuits[pB]
		delete(circuits, pB)
		circuits[pA] = append(circuits[pA], extra...)
		for _, C := range extra {
			roots[C] = pA
		}
		if len(circuits[pA]) >= len(boxes) {
			found = &c
			break
		}
	}

	if found == nil {
		return fmt.Errorf("uncomplitable circuit")
	}

	fmt.Fprintln(w, found.A.X*found.B.X)

	return nil
}
