package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "input.txt"

func solve() (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var checksum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		if len(values) == 0 {
			continue
		}

		converted := make([]int, len(values))
		for index, val := range values {
			i, err := strconv.Atoi(val)
			if err != nil {
				return -1, err
			}

			converted[index] = i
		}

		for i := 0; i < len(converted); i++ {
			for j := i + 1; j < len(converted); j++ {
				fmt.Println(converted[i], converted[j])

				if converted[i]%converted[j] == 0 {
					checksum += converted[i] / converted[j]
					break
				} else if converted[j]%converted[i] == 0 {
					checksum += converted[j] / converted[i]
					break
				}
			}
		}
	}

	return checksum, scanner.Err()
}

func main() {
	checksum, err := solve()
	if err != nil {
		panic(err)
	}

	fmt.Println(checksum)
}
