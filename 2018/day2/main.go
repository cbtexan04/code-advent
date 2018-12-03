package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const inputPath = "input.txt"

func solve(input io.Reader) int {
	var twoCount, threeCount int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var foundTwo, foundThree bool
		line := scanner.Text()

		for _, l := range line {
			letter := string(l)
			count := strings.Count(line, letter)
			if count == 2 && !foundTwo {
				foundTwo = true
				twoCount++
			} else if count == 3 && !foundThree {
				foundThree = true
				threeCount++
			}

			if foundTwo && foundThree {
				break
			}
		}
	}

	return twoCount * threeCount
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := solve(file)
	fmt.Println(result)
}
