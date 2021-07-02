package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type OpenFileError struct {
	err  error
	file string
}

type CantReadFileError struct {
	err  error
	file string
}

func (e OpenFileError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("Error during file %q opening: %v", e.file, e.err.Error())
	}
	return fmt.Sprintf("Unknown error during file %q opening", e.file)
}

func (e CantReadFileError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("Error during file %q reading: %v", e.file, e.err.Error())
	}
	return fmt.Sprintf("Unknown error during file %q reading", e.file)
}

func readFile(path string) ([]byte, error) {
	fd, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, OpenFileError{err, path}
	}
	raw, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, CantReadFileError{err, path}
	}
	return raw, nil
}

func countLines(raw []byte) int {
	return bytes.Count(raw, []byte("\n")) + 1 // because file without any '\n' sign has already one line
}

func main() {
	path := "./README.md"
	raw, err := readFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(countLines(raw))
}
