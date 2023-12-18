package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func ErrorReadingFile(fileName string) error {
	return errors.New(fmt.Sprintf("Error reading file %s\n", fileName))
}

func ErrorReadingStdin(err error) error {
	return errors.New(fmt.Sprintf("Error reading stdin %s\n", err))
}

func NoOfBytes(content []byte) int {
	return len(content)
}

func NoOfLines(content []byte) int {
	return bytes.Count(content, []byte("\n"))
}

func NoOfWords(content []byte) int {
	return len(bytes.Fields(content))
}

func NoOfChars(content []byte) int {
	return len(bytes.Runes(content))
}

func ReadStdin() ([]byte, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, ErrorReadingStdin(err)
	}

	return b, nil
}

func ReadFile(fileName string) ([]byte, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, ErrorReadingFile(fileName)
	}

	return b, nil
}
