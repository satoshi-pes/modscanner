package modscanner

import (
	"bufio"
	"io"
)

type Scanner struct {
	// current line number
	currentLineNum int

	r *io.Reader
	s *bufio.Scanner
}

// NewScanner returns a new Scanner to read from r, but unlike the bufio.scanner
// in the standard library, it retains the *io.Reader and the currentline number.
func NewScanner(r io.Reader) *Scanner {
	var s Scanner
	s.r = &r
	s.s = bufio.NewScanner(r)

	return &s
}

// Scan advances the scanner to the next token.
// This is a wrapper of bufio.scanner.Scan().
func (s *Scanner) Scan() bool {
	ok := s.s.Scan()
	if ok {
		s.currentLineNum++
	}
	return ok
}

// LineNumber returns the current line number of the file.
func (s *Scanner) LineNumber() int {
	return s.currentLineNum
}

// Buffer is a wrapper of bufio.scanner.Buffer().
func (s *Scanner) Buffer(buf []byte, max int) {
	s.s.Buffer(buf, max)
}

// Bytes is a wrapper of bufio.scanner.Bytes().
func (s *Scanner) Bytes() []byte {
	return s.s.Bytes()
}

// Err is a wrapper of bufio.scanner.Err().
func (s *Scanner) Err() error {
	return s.s.Err()
}

// Split is a wrapper of bufio.scanner.Split().
func (s *Scanner) Split(split bufio.SplitFunc) {
	s.s.Split(split)
}

// Text returns the most recent token.
// This is a wrapper of bufio.scanner.Text().
func (s *Scanner) Text() string {
	return s.s.Text()
}
