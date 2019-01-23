package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func args() ([]string, string, error) {
	if len(os.Args) != 4 || os.Args[1] != "-f" {
		fmt.Println("usage: fgrep-go -f words.txt path")

		return nil, "", errors.New("usage error")
	}

	content, err := ioutil.ReadFile(os.Args[2])

	if err != nil {
		fmt.Println("error loading words:", err)
		return nil, "", err
	}

	words := strings.Split(string(content), "\n")
	words = words[:len(words) - 1]

	return words, os.Args[3], nil
}

func matchInRed(line string, matches []Match) string {
	red :="\033[0;31m"
	noColor :="\033[0m"

	result := ""
	currentEnd := 0

	for _, match := range matches {
		if match.end <= currentEnd {
			continue
		}

		if match.start <= currentEnd {
			result += red + line[currentEnd: match.end] + noColor
		} else {
			result += line[currentEnd: match.start] + red + match.word + noColor
		}

		currentEnd = match.end
	}

	return result + line[currentEnd:]
}

func searchInFile(path string, words []string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error reading:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := Search(line, words...)

		if len(matches) > 0 {
			fmt.Println(matchInRed(line, matches))
		}
	}
}

func main() {
	words, path, err := args()

	if err != nil {
		return
	}

	searchInFile(path, words)
}
