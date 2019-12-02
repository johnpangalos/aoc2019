package fileparse

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ScannerInterface describes
type ScannerInterface interface {
	CommaStringParse(string) []string
}

// Scanner extends the bufio Scanner struct and adds functions on
// it throught ScannerInterface
type Scanner struct {
	*bufio.Scanner
}

// NewScanner reads the input var @path and returns a bufio scanner
func NewScanner(path string) (*bufio.Scanner, *os.File, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	return scanner, file, nil
}

// CommaStringParse returns a array of strings from a @str
// splitting them on commas
func (s *Scanner) CommaStringParse() []string {
	return strings.Split(s.Text(), ",")
}
