package main

import (
	"log"
	"strconv"
)

func main() {
	// Update input to update advent's input
	input := "123"

	sum, err := solve(input)
	if err != nil {
		panic(err)
	}

	log.Println(sum)
}

func solve(input string) (int, error) {
	length := len(input)
	var sum int

	for i := 0; i < len(input); i++ {
		newIndex := getNewIndex(i, length)
		y := string(input[newIndex])
		x := string(input[i])

		if x == y {
			integer, err := strconv.Atoi(x)
			if err != nil {
				return sum, err
			}

			sum += integer
		}
	}

	return sum, nil
}

func getNewIndex(startIndex int, length int) int {
	index := startIndex + (length / 2)
	if index >= length {
		index = index - length
	}

	return index
}
