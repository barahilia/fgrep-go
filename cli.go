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
			fmt.Println(line)

			for _, match := range matches {
				fmt.Println("  -", match.word)
			}
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
