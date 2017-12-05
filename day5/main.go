package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const input = "input.txt"

func main() {
	arr, err := loadSteps()
	if err != nil {
		panic(err)
	}

	steps := solve(arr)
	fmt.Println("Number of steps:", steps)
}

func loadSteps() ([]int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		lines = append(lines, n)
	}

	return lines, scanner.Err()
}

func solve(arr []int) int {
	var index, steps int

	for index < len(arr) && index >= 0 {
		element := arr[index]
		if element >= 3 {
			arr[index] = element - 1
		} else {
			arr[index] = element + 1
		}

		index = (index + element)
		steps++
	}

	return steps
}
