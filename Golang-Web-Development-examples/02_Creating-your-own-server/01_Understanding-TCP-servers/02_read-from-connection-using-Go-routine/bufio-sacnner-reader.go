package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Honesty is the best policy.\n Man you in to change your policy now.\n Because this is the truth of the world."

	// func NewScanner(r io.Reader) *Scanner
	// NewScanner takes reader (r io.reader)
	// returns *Scanner
	scanner := bufio.NewScanner(strings.NewReader(s))

	//scanner.Split(bufio.ScanWords)

	// func (s *Scanner) Scan() bool
	// func Scan returns boolean value
	for scanner.Scan() {
		// func (s *Scanner) Text() string
		// returns string
		line := scanner.Text()
		fmt.Println(line)
	}
}
