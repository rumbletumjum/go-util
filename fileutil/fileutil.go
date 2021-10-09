// Good artists create, great artists steal.
// Seriously though, writing just-barely-different versions of stdlib funcs is a
// good way to get good idioms under your fingers. For me, at least.
package fileutil

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return lines, err
}

// lazy person's WriteFile
func WriteFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
