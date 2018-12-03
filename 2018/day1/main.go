package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const inputPath = "input.txt"

func readerToInputs(input io.Reader) ([]int, error) {
	var newInputs []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "+")

		i, err := strconv.Atoi(line)
		if err != nil {
			return newInputs, err
		}

		newInputs = append(newInputs, i)
	}

	return newInputs, nil
}

func solvePart1(inputs []int) (int, error) {
	var sum int
	for _, i := range inputs {
		sum = sum + i
	}

	return sum, nil
}

func solvePart2(inputs []int) int {
	seenMap := make(map[int]bool)

	var sum int

Beginning:
	for index, i := range inputs {
		sum = sum + i

		if _, ok := seenMap[sum]; ok {
			return sum
		} else {
			seenMap[sum] = true
		}

		// Last one
		if index == len(inputs)-1 {
			goto Beginning
		}
	}

	return sum
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inputs, err := readerToInputs(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePart1(inputs))
	fmt.Println(solvePart2(inputs))
}
