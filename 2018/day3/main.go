package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const inputPath = "input.txt"

var ClaimRE = regexp.MustCompile("([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")

type Claim struct {
	XPos   int
	YPos   int
	Height int
	Width  int
	ID     int
}

type Grid [][]int

func NewClaim(s string) (Claim, error) {
	matches := ClaimRE.FindStringSubmatch(s)
	if len(matches) != 6 {
		fmt.Println(s, len(matches))
		return Claim{}, errors.New("invalid line")
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return Claim{}, err
	}

	ypos, err := strconv.Atoi(matches[3])
	if err != nil {
		return Claim{}, err
	}

	xpos, err := strconv.Atoi(matches[2])
	if err != nil {
		return Claim{}, err
	}

	width, err := strconv.Atoi(matches[4])
	if err != nil {
		return Claim{}, err
	}

	height, err := strconv.Atoi(matches[5])
	if err != nil {
		return Claim{}, err
	}

	return Claim{
		ID:     id,
		YPos:   ypos,
		XPos:   xpos,
		Width:  width,
		Height: height,
	}, nil
}

func gatherInputs(input io.Reader) ([]Claim, error) {
	var claims []Claim

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		c, err := NewClaim(line)
		if err != nil {
			return claims, err
		}

		claims = append(claims, c)
	}

	return claims, nil
}

func solve(claims []Claim) (int, error) {
	var grid Grid

	for _, c := range claims {
		grid = grid.AddClaim(c)
	}

	return grid.NumOverlaps(), nil
}

func (g Grid) NumOverlaps() int {
	var count int
	for _, row := range g {
		for _, id := range row {
			if id == -1 {
				count++
			}
		}
	}

	return count
}

func (g Grid) AddClaim(c Claim) Grid {
	numRowsNeeded := c.Height + c.YPos
	minRowWidth := c.Width + c.XPos + 1

	// Create the number of rows needed
	if len(g) < numRowsNeeded {
		for i := len(g); i < numRowsNeeded+1; i++ {
			newRow := make([]int, minRowWidth)
			g = append(g, newRow)
		}
	}

	// Make sure we have our min width
	for i, row := range g {
		if len(row) < minRowWidth {
			extraIndices := make([]int, minRowWidth-len(row))
			g[i] = append(g[i], extraIndices...)
		}
	}

	// Mark the appropriate grid indexes
	for i := c.YPos; i < c.YPos+c.Height; i++ {
		for j := c.XPos; j < c.XPos+c.Width; j++ {
			if g[i][j] == 0 {
				g[i][j] = c.ID
			} else {
				g[i][j] = -1
			}
		}
	}

	return g
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	claims, err := gatherInputs(file)
	if err != nil {
		panic(err)
	}

	result, err := solve(claims)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
