package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var input string

	switch len(os.Args) {
	case 1:
		bs, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Reading stdin: %s\n", err)
			os.Exit(1)
		}
		input = string(bs)
	case 2:
		bs, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Opening %s: %s\n", os.Args[1], err)
			os.Exit(1)
		}
		input = string(bs)
	default:
		fmt.Fprintf(os.Stderr, "Usage: `%[1]s file.mms` or `cat file.mms | %[1]s`\n", os.Args[0])
		os.Exit(1)
	}

	var lines []line
	for _, str := range strings.Split(input, "\n") {
		lines = append(lines, readLine(str))
	}

	printLines(lines)
}
