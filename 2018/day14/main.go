package main

import (
	"fmt"
	"strconv"
)

type RecipeBoard struct {
	Iterations int
	Numbers    []int
	ElfAIndex  int
	ElfBIndex  int
}

func main() {
	board := RecipeBoard{
		Iterations: 1,
		Numbers:    []int{3, 7},
		ElfAIndex:  0,
		ElfBIndex:  1,
	}

	input := 540561
	for i := 0; i < input+10; i++ {
		board.Increment()
	}

	fmt.Println(board.Numbers[input : input+10])
}

func (b *RecipeBoard) Increment() {
	elfANum := b.Numbers[b.ElfAIndex]
	elfBNum := b.Numbers[b.ElfBIndex]

	for _, r := range fmt.Sprintf("%d", elfANum+elfBNum) {
		n, _ := strconv.Atoi(string(r))
		b.Numbers = append(b.Numbers, n)
	}

	b.ElfAIndex = (elfANum + b.ElfAIndex + 1) % (len(b.Numbers))
	b.ElfBIndex = (elfBNum + b.ElfBIndex + 1) % (len(b.Numbers))

	b.Iterations++
}

func (b *RecipeBoard) Print() {
	for i, n := range b.Numbers {
		if i == b.ElfBIndex {
			fmt.Printf(" [%d] ", n)
		} else if i == b.ElfAIndex {
			fmt.Printf(" (%d) ", n)
		} else {
			fmt.Printf(" %d ", n)
		}
	}

	fmt.Println()
}
