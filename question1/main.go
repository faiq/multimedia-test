package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}

//go through every line and
func processLines(lines []string) {
	for _, line := range lines {
		stop := strings.IndexAny(line, "!?.")
		if stop > -1 {
			line = line[:stop]
		}
		fmt.Printf("this is line %s \n", line)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("looks like you didn't insert a text file to analyze")
		return
	}

	fileToRead := os.Args[1]
	lines, err := readLines(fileToRead)
	if err != nil {
		fmt.Printf("looks like there was an error reading the file")
		return
	}
	processLines(lines)
}
