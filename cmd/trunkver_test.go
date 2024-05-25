package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassAllExplicitArgs(t *testing.T) {
	args := []string{
		"trunkver",
		"--timestamp", "2024-05-22T16:25:48+02:00",
		"--source-ref", "gb4dc0d3",
		"--build-ref", "12345",
	}
	readStdout, writeStdout, _ := os.Pipe()
	_, writeStderr, _ := os.Pipe()
	run(writeStdout, writeStderr, args)
	writeStdout.Close()
	writeStderr.Close()

	out, _ := io.ReadAll(readStdout)
	assert.Equal(t,
		"20240522142548.0.0-gb4dc0d3-12345\n",
		string(out),
		"Should print the version to stdout",
	)
}

func TestPrintsVersion(t *testing.T) {
	args := []string{
		"trunkver",
		"--version",
	}
	_, writeStdout, _ := os.Pipe()
	readStderr, writeStderr, _ := os.Pipe()
	run(writeStdout, writeStderr, args)
	writeStdout.Close()
	writeStderr.Close()

	err, _ := io.ReadAll(readStderr)
	assert.Equal(t,
		"0.0.0-HEAD-local\n",
		string(err),
		"Should print the version to stdout",
	)
}
