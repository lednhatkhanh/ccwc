package main

import (
	"bytes"
	"ccwc/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BytesCount(t *testing.T) {
	actual := bytes.NewBufferString("")
	cmd := cmd.NewRootCmd()
	cmd.SetOut(actual)
	cmd.SetArgs([]string{"-c", "test.txt"})
	cmd.Execute()

	expected := "335045 test.txt\n"

	assert.Equal(t, expected, actual.String())
}

func Test_LinesCount(t *testing.T) {
	actual := bytes.NewBufferString("")
	cmd := cmd.NewRootCmd()
	cmd.SetOut(actual)
	cmd.SetArgs([]string{"-l", "test.txt"})
	cmd.Execute()

	expected := "7145 test.txt\n"

	assert.Equal(t, expected, actual.String())
}

func Test_WordsCount(t *testing.T) {
	actual := bytes.NewBufferString("")
	cmd := cmd.NewRootCmd()
	cmd.SetOut(actual)
	cmd.SetArgs([]string{"-w", "test.txt"})
	cmd.Execute()

	expected := "58164 test.txt\n"

	assert.Equal(t, expected, actual.String())
}

func Test_CharactersCount(t *testing.T) {
	actual := bytes.NewBufferString("")
	cmd := cmd.NewRootCmd()
	cmd.SetOut(actual)
	cmd.SetArgs([]string{"-m", "test.txt"})
	cmd.Execute()

	expected := "332147 test.txt\n"

	assert.Equal(t, expected, actual.String())
}

func Test_WithoutFlags(t *testing.T) {
	actual := bytes.NewBufferString("")
	cmd := cmd.NewRootCmd()
	cmd.SetOut(actual)
	cmd.SetArgs([]string{"test.txt"})
	cmd.Execute()

	expected := "7145 58164 335045 test.txt\n"

	assert.Equal(t, expected, actual.String())
}
