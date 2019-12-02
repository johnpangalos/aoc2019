package fileparse

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Scanner extends the bufio Scanner struct and adds functions on
// it throught ScannerInterface
type Scanner struct {
	*bufio.Scanner
	file *os.File
}

// NewScanner reads the input var @path and returns a bufio scanner
func NewScanner(path string) (*Scanner, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	s := &Scanner{
		scanner,
		file,
	}
	return s, nil
}

// CommaStringParse returns a array of strings from a @str
// splitting them on commas
func (s *Scanner) CommaStringParse() []string {
	s.Scan()
	return strings.Split(s.Scanner.Text(), ",")
}

// CommaStringParseInt returns a array of integers from a @str
// splitting them on commas
func (s *Scanner) CommaStringParseInt() []int64 {
	strs := s.CommaStringParse()
	var ints []int64
	for _, val := range strs {
		valInt, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ints = append(ints, valInt)
	}
	return ints
}

// Close closes the file attached to the scanner
func (s *Scanner) Close() {
	s.file.Close()
}
