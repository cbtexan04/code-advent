package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

const inputPath = "input.txt"

var NodeRE = regexp.MustCompile("Step ([a-zA-Z]) must be finished before step ([a-zA-Z]) can begin")

func solve(m map[rune]map[rune]bool) (string, error) {
	steps := []string{}

	for len(m) > 0 {
		var availableNodes []string

		// Find our nodes which don't have any dependencies
		for r, deps := range m {
			if len(deps) == 0 {
				availableNodes = append(availableNodes, string(r))
			}
		}

		sort.Slice(availableNodes, func(i, j int) bool {
			return rune(availableNodes[i][0]) < rune(availableNodes[j][0])
		})

		if len(availableNodes) == 0 {
			return "", errors.New("unable to solve puzzle")
		}

		// Our first node can be used as the next step
		nextStep := availableNodes[0]
		steps = append(steps, nextStep)

		// We need to remove this dep from any other node
		for key, deps := range m {
			for k, _ := range deps {
				if string(k) == nextStep {
					delete(m[key], k)
				}
			}
		}

		// Also remove it from the top level map so we don't do it again
		delete(m, rune(availableNodes[0][0]))
	}

	return strings.Join(steps, ""), nil
}

func gatherInputs(input io.Reader) (map[rune]map[rune]bool, error) {
	m := make(map[rune]map[rune]bool)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		matches := NodeRE.FindStringSubmatch(line)
		if len(matches) != 3 || len(matches[1]) != 1 || len(matches[2]) != 1 {
			return m, errors.New("invalid line")
		}

		dependency := rune(matches[1][0])
		node := rune(matches[2][0])

		if _, ok := m[dependency]; !ok {
			m[dependency] = make(map[rune]bool)
		}

		if _, ok := m[node]; !ok {
			m[node] = make(map[rune]bool)
		}

		m[node][dependency] = true
	}

	return m, nil
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m, err := gatherInputs(file)
	if err != nil {
		panic(err)
	}

	result, err := solve(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
