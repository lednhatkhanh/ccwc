package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	bytesFlag := flag.Bool("c", false, "Outputs the number of bytes in a file.")
	linesFlag := flag.Bool("l", false, "Outputs the number of lines in a file.")
	wordsFlag := flag.Bool("w", false, "Outputs the number of words in a file.")
	charsFlag := flag.Bool("m", false, "Outputs the number of characters in a file.")

	flag.Parse()

	args := flag.Args()
	var fileName string
	var fileContent []byte

	if len(args) == 1 {
		filename := args[0]
		content, err := readFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fileName = filename
		fileContent = content
	} else {
		stdin, err := readStdin()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fileContent = stdin
	}

	output := ""

	switch {
	case *bytesFlag:
		output += fmt.Sprintf("%d", noOfBytes(fileContent))
	case *linesFlag:
		output += fmt.Sprintf("%d", noOfLines(fileContent))
	case *wordsFlag:
		output += fmt.Sprintf("%d", noOfWords(fileContent))
	case *charsFlag:
		output += fmt.Sprintf("%d", noOfChars(fileContent))
	default:
		output += fmt.Sprintf("%d %d %d", noOfLines(fileContent), noOfWords(fileContent), noOfBytes(fileContent))
	}

	if fileName != "" {
		output += fmt.Sprintf(" %s\n", fileName)
	} else {
		output += "\n"
	}

	fmt.Printf(output)
}

func noOfBytes(content []byte) int {
	return len(content)
}

func noOfLines(content []byte) int {
	return bytes.Count(content, []byte("\n"))
}

func noOfWords(content []byte) int {
	return len(bytes.Fields(content))
}

func noOfChars(content []byte) int {
	return len(bytes.Runes(content))
}

func readStdin() ([]byte, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading stdin %s\n", err))
	}

	return b, nil
}

func readFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file %s\n", filename))
	}

	return b, nil
}
