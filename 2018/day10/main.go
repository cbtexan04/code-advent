package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const solutionFile = "output.txt"
const inputPath = "input.txt"

var NodeRE = regexp.MustCompile("position=<\\s?([-0-9]+),\\s?\\s?([-0-9]+)> velocity=<\\s?([-0-9]+),\\s*?([-0-9]+)>")

type Mark struct {
	CurrentX  int
	CurrentY  int
	VelocityX int
	VelocityY int
}

type Solver struct {
	Board [][]string
	Marks []Mark
}

func solve(s Solver) error {
	for times := 0; times < 3; times++ {
		for i, m := range s.Marks {
			// Remember our old positions
			oldx := s.Marks[i].CurrentX
			oldy := s.Marks[i].CurrentY

			// Mark our next spot
			s.Marks[i].CurrentX += m.VelocityX
			s.Marks[i].CurrentY += m.VelocityY

			// Need to know if we have to increase the board
			minWidthNeeded := s.Marks[i].CurrentX
			minRowsNeeded := s.Marks[i].CurrentY

			// TODO: so we have an issue here.. we could totally
			// have a negative velocity or negative starting point,
			// which could cause an OOB exception. This will need
			// some work in order to normalize our marks so that
			// the grid works in all directions.

			// Create the number of rows needed
			if len(s.Board) < minRowsNeeded {
				for i := len(s.Board); i < minRowsNeeded+1; i++ {
					newRow := make([]string, minWidthNeeded)
					// Fill out the default spaces
					for j, _ := range newRow {
						newRow[j] = "."
					}

					s.Board = append(s.Board, newRow)
				}
			}

			// Make sure we have our min width
			for i, row := range s.Board {
				if len(row) < minWidthNeeded {
					extraIndices := make([]string, minRowsNeeded-len(row))
					// Fill out the default spaces
					for j, _ := range extraIndices {
						extraIndices[j] = "."
					}
					s.Board[i] = append(s.Board[i], extraIndices...)
				}
			}

			// Mark our old position as not being used anymore, and
			// our new as being utilized
			s.Board[oldy][oldx-1] = "."
			s.Board[s.Marks[i].CurrentY][s.Marks[i].CurrentX-1] = "#"
		}

		s.Print()
	}

	return nil
}

func (s Solver) Print() {
	for _, row := range s.Board {
		fmt.Println(row)
	}

	fmt.Printf("\n--------------\n")
}

func gatherInputs(input io.Reader) (Solver, error) {
	var s Solver

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		matches := NodeRE.FindStringSubmatch(line)
		if len(matches) != 5 {
			return s, errors.New("invalid input")
		}

		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		xvol, _ := strconv.Atoi(matches[3])
		yvol, _ := strconv.Atoi(matches[4])

		m := Mark{
			CurrentX:  x,
			CurrentY:  y,
			VelocityX: xvol,
			VelocityY: yvol,
		}

		s.Marks = append(s.Marks, m)
	}

	return s, nil
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s, err := gatherInputs(file)
	if err != nil {
		panic(err)
	}

	solve(s)
}
