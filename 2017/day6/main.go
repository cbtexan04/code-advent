package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "input.txt"

type Solver struct {
	Initial   []int
	Solutions [][]int
}

func (s *Solver) solve() {
	arr := s.Initial

	for {
		if !s.hasSeen(arr) {
			s.Solutions = append(s.Solutions, arr)
		} else {
			fmt.Println(s.Solutions, arr)
			break
		}

		// Make a copy, otherwise we're changing the value for all the other
		// solution indexes
		arr = append([]int(nil), arr...)

		// Get the biggest value/index, and reset it to 0
		biggestIndex := getLargestIndex(arr)
		biggestValue := arr[biggestIndex]
		arr[biggestIndex] = 0

		// Start at the next element
		currentIndex := biggestIndex + 1

		// Traverse the array until our biggestValue has been drained
		for biggestValue > 0 {
			if currentIndex == len(arr) {
				currentIndex = 0
			}

			arr[currentIndex] = arr[currentIndex] + 1
			currentIndex++
			biggestValue--
		}
	}
}

func (s *Solver) hasSeen(arr []int) bool {
	for i := 0; i < len(s.Solutions); i++ {
		if equalSlices(arr, s.Solutions[i]) {
			return true
		}

	}

	return false

}

func equalSlices(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func getInput() ([]int, error) {
	puzzleInput := make([]int, 0)

	file, err := os.Open(input)
	if err != nil {
		return puzzleInput, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		for _, v := range values {
			i, err := strconv.Atoi(v)
			if err != nil {
				return puzzleInput, err
			}

			puzzleInput = append(puzzleInput, i)
		}
	}

	return puzzleInput, scanner.Err()
}

func getLargestIndex(arr []int) int {
	var largestIndex int
	var largestNum int

	for i, v := range arr {
		if v > largestNum {
			largestNum = v
			largestIndex = i
		}
	}

	return largestIndex
}

func main() {
	arr, err := getInput()
	if err != nil {
		panic(err)
	}

	s := &Solver{
		Initial:   arr,
		Solutions: make([][]int, 0),
	}

	s.solve()
	fmt.Println(len(s.Solutions))
}
