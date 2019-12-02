package fileparse

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ScannerInterface describes
type ScannerInterface interface {
	CommaStringParse(string) []string
}

// ScanWrap extends the bufio Scanner struct and adds functions on
// it throught ScannerInterface
type ScanWrap struct {
	Scanner *bufio.Scanner
}

// NewScanner reads the input var @path and returns a bufio scanner
func NewScanner(path string) (*ScanWrap, *os.File, error) {
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
	s := &ScanWrap{
		Scanner: scanner,
	}
	return s, file, nil
}

// CommaStringParse returns a array of strings from a @str
// splitting them on commas
func (s *ScanWrap) CommaStringParse() []string {
	s.Scanner.Scan()
	return strings.Split(s.Scanner.Text(), ",")
}

// CommaStringParseInt returns a array of integers from a @str
// splitting them on commas
func (s *ScanWrap) CommaStringParseInt() []int64 {
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
