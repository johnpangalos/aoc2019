package fileparse

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// NewScanner reads the input var PATH and returns a bufio scanner
func NewScanner(path string) (*bufio.Scanner, error) {
	path, err := filepath.Abs(path)
	fmt.Println(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return scanner, nil
}
