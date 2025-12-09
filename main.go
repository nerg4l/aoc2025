package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"text/template"
)

func main() {
	flag.Parse()

	var d Day

	switch flag.Arg(0) {
	case "add":
		addDay(flag.Arg(1))
		return
	case "1", "01":
		d = Day01{}
	case "2", "02":
		d = Day02{}
	case "3", "03":
		d = Day03{}
	case "4", "04":
		d = Day04{}
	case "5", "05":
		d = Day05{}
	case "6", "06":
		d = Day06{}
	default:
		fmt.Fprintf(os.Stderr, "Invalid day argument\n")
		os.Exit(1)
	}

	var f func(io.Reader, io.Writer) error = nil
	switch flag.Arg(1) {
	case "1":
		f = d.Segment1
	case "2":
		f = d.Segment2
	default:
		fmt.Fprintf(os.Stderr, "Invalid segment argument\n")
		os.Exit(1)
	}

	if err := f(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

type Day interface {
	Segment1(r io.Reader, w io.Writer) error
	Segment2(r io.Reader, w io.Writer) error
}

func addDay(day string) {
	if len(day) == 0 {
		fmt.Fprintf(os.Stderr, "Invalid day argument\n")
		os.Exit(1)
	}
	if d, err := strconv.Atoi(day); err != nil || d < 0 {
		fmt.Fprintf(os.Stderr, "Invalid day argument\n")
		os.Exit(1)
	}

	fname := fmt.Sprintf("./day%s.go", day)
	if _, err := os.Stat(fname); !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "%s already exists\n", fname)
		os.Exit(1)
	}
	fnameTest := fmt.Sprintf("./day%s_test.go", day)
	if _, err := os.Stat(fname); !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "%s already exists\n", fnameTest)
		os.Exit(1)
	}

	tmpl := template.Must(template.New("").Parse(`package main

import (
	"bufio"
	"fmt"
	"io"
)

type Day{{.}} struct{}

func (Day{{.}}) Segment1(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		s.Text()
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w)

	return nil
}

func (Day{{.}}) Segment2(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		s.Text()
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Fprintln(w)

	return nil
}
`))
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open %q: %s", fname, err))
	}
	tmpl.Execute(f, day)

	tmplTest := template.Must(template.New("").Parse(`package main

import (
	"io"
	"strings"
	"testing"
)

func TestDay{{.}}(t *testing.T) {
	tests := []struct {
		name  string
		input io.ReadSeeker
		want1 string
		want2 string
	}{
		{
			name: "small",
			input: strings.NewReader(` + "`" + `
` + "`" + `),
			want1: ` + "`" + `
` + "`" + `,
			want2: ` + "`" + `
` + "`" + `,
		},
		{
			name:  "big",
			input: mayOpen("./day{{.}}.input"),
			want1: ` + "`" + `
` + "`" + `,
			want2: ` + "`" + `
` + "`" + `,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, dayTester(tt, Day{{.}}{}))
	}
}
`))
	f, err = os.OpenFile(fnameTest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open %q: %s", fname, err))
	}
	tmplTest.Execute(f, day)
}
