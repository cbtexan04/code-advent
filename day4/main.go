package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const input = "input.txt"

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func sortString(w string) string {
	r := []rune(w)
	sort.Sort(runes(r))
	return string(r)
}

func areAnagrams(str1, str2 string) bool {
	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)
	return sortedStr1 == sortedStr2
}

func containsAnagrams(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if areAnagrams(arr[i], arr[j]) {
				return true
			}
		}
	}

	return false
}

func solve() (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		if !containsAnagrams(values) {
			sum += 1
		}
	}

	return sum, nil
}

func main() {
	solution, err := solve()
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
