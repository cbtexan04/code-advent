package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

const inputPath = "input.txt"

func solvePart1(input []rune) (int, error) {
	for true {
		for i, r := range input {
			if i == len(input)-1 {
				return len(input), nil
			}

			if isPolyMatch(r, input[i+1]) {
				// Remove the poly match from the array and go
				// back through
				input = append(input[:i], input[i+2:]...)
				break
			}
		}
	}

	return -1, nil
}

func isPolyMatch(r1 rune, r2 rune) bool {
	if unicode.IsUpper(r1) && unicode.IsUpper(r2) {
		// No match if both are uppercased
		return false
	} else if !unicode.IsUpper(r1) && !unicode.IsUpper(r2) {
		// No match if both are lowercased
		return false
	}

	// One is uppercased and one is lowercased. We can convert them both to
	// uppercase and check for equality
	return unicode.ToUpper(r1) == unicode.ToUpper(r2)
}

func main() {
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	b = bytes.TrimSpace(b)

	// Convert to runes to make life easier
	input := make([]rune, len(b))
	for i, r := range b {
		input[i] = rune(r)
	}

	fmt.Println(solvePart1(input))
}
