package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

	state := make(map[string]int) //state will keep a mapping of sanitized names

	for _, line := range lines {
		line = strings.ToLower(line)
		stop := strings.IndexAny(line, "!?.\n")
		if stop > -1 {
			line = line[:stop]
		}
		if state[line] != 0 {
			state[line] = state[line] + 1
		} else {
			state[line] = 1
		}
	}
	pl := make(PairList, len(lines))
	i := 0
	for k, v := range state {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	for _, pair := range pl {
		if pair.Value > 1 {
			fmt.Printf("%s %d\n", pair.Key, pair.Value)
		}
	}
}
