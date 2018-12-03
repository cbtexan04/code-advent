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

	result, err := solvePart1(inputs)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
