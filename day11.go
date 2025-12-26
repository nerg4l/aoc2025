package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strings"
)

type Day11 struct{}

type DeviceConnection struct {
	input   string
	outputs []string
}

func (d Day11) Lines(s *bufio.Scanner) iter.Seq2[DeviceConnection, error] {
	return func(yield func(DeviceConnection, error) bool) {
		for s.Scan() {
			str := strings.Split(s.Text(), " ")
			conn := DeviceConnection{input: str[0], outputs: str[1:]}
			if conn.input[len(conn.input)-1] != ':' {
				yield(DeviceConnection{}, fmt.Errorf("invalid input %q", conn.input))
				return
			}
			conn.input = conn.input[:len(conn.input)-1]
			if !yield(conn, nil) {
				return
			}
		}
	}
}

func (d Day11) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	conns := map[string]DeviceConnection{}

	for conn, err := range d.Lines(s) {
		if err != nil {
			return err
		}
		conns[conn.input] = conn
	}

	if err := s.Err(); err != nil {
		return err
	}

	result := 0
	if me, ok := conns["you"]; ok {
		result = CountPathToNode(conns, me.outputs, "out")
	}
	fmt.Fprintln(w, result)

	return nil
}

func (d Day11) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	conns := map[string]DeviceConnection{}

	for conn, err := range d.Lines(s) {
		if err != nil {
			return err
		}
		conns[conn.input] = conn
	}

	if err := s.Err(); err != nil {
		return err
	}

	result := 0
	if svr, ok := conns["svr"]; ok {
		result += CountPathToNode(conns, svr.outputs, "fft") *
			CountPathToNode(conns, conns["fft"].outputs, "dac") *
			CountPathToNode(conns, conns["dac"].outputs, "out")
		result += CountPathToNode(conns, svr.outputs, "dac") *
			CountPathToNode(conns, conns["dac"].outputs, "fft") *
			CountPathToNode(conns, conns["fft"].outputs, "out")
	}
	fmt.Fprintln(w, result)

	return nil
}

func CountPathToNode(conns map[string]DeviceConnection, outputs []string, end string) int {
	mem := map[string]int{}
	return listPathToOutput(conns, outputs, end, mem, 0)
}

func listPathToOutput(conns map[string]DeviceConnection, outputs []string, end string, mem map[string]int, depth int) int {
	if depth >= len(conns) {
		panic("recursive data")
	}
	var result int
	for _, out := range outputs {
		if out == end {
			result++
			continue
		}
		if n, ok := mem[out]; ok {
			result += n
			continue
		}
		n := listPathToOutput(conns, conns[out].outputs, end, mem, depth+1)
		mem[out] = n
		result += n
	}
	return result
}
