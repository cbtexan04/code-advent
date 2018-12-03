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

func solvePart1(input io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "+")

		i, err := strconv.Atoi(line)
		if err != nil {
			return i, err
		}

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

	result, err := solve(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
