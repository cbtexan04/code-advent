package main

import "fmt"

const input = 1718
const gridSize = 3

func calcPower(x int, y int) int {
	id := x + 10
	p := ((x + 10) * y) + input
	return ((p * id / 100) % 10) - 5
}

func solve() (xcoord int, ycoord int) {
	var highestPowerSeen int

	for x := 0; x < 300-gridSize; x++ {
		for y := 0; y < 300-gridSize; y++ {
			var power int
			for i := 0; i < gridSize; i++ {
				for j := 0; j < gridSize; j++ {
					power += calcPower(x+i, y+j)
				}
			}

			if power > highestPowerSeen {
				xcoord = x
				ycoord = y
				highestPowerSeen = power
			}
		}
	}

	return xcoord, ycoord
}

func main() {
	fmt.Println(solve())
}
