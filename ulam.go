package main

import (
	"flag"
	"fmt"
	"os"
)

var spiralDim int
var spiralStart int

const (
	right = iota
	down  = iota
	left  = iota
	up    = iota
)

func printUsage() {
	fmt.Println("Usage:\n\tulam -d <dimension>")
}

func init() {
	flag.IntVar(&spiralDim, "d", 0, "Spiral dimension")
	flag.IntVar(&spiralStart, "s", 1, "Spiral starting value")
	if spiralStart < 1 {
		spiralStart = 1
	}
}

func makeUlamArray(dim int, val int) []int {
	ulam := make([]int, dim*dim)

	var row int
	var col int
	var lastCell int
	currSqBase := 1
	placed := 0
	totalPlaced := 0
	direction := right

	if (dim % 2) != 0 {
		row = dim / 2
		col = dim / 2
		lastCell = dim - 1
	} else {
		row = dim/2 - 1
		col = dim/2 - 1
		lastCell = row * (dim - 1)
	}

	for running := true; running; {
		if ((row * dim) + col) != lastCell {
			running = false
		}
		ulam[row*dim+col] = val
		val++
		placed++
		totalPlaced++

		if (totalPlaced != (currSqBase * currSqBase)) || (!running) {
			if placed == currSqBase {
				switch direction {
				case right:
					direction = down
				case down:
					direction = left
				case left:
					direction = up
				case up:
					direction = right
				}
				placed = 0
			}
		} else {
			switch direction {
			case right:
				col++
			case down:
				row++
			case left:
				col--
			case up:
				row--
			}
			ulam[row*dim+col] = val
			val++
			placed++
			totalPlaced++
			switch direction {
			case right:
				direction = down
			case down:
				direction = left
			case left:
				direction = up
			case up:
				direction = right
			}
			currSqBase++
			placed = 0
		}
		switch direction {
		case right:
			col++
		case down:
			row++
		case left:
			col--
		case up:
			row--
		}
	}
	return ulam
}

func main() {
	flag.Parse()
	if spiralDim < 1 {
		printUsage()
		os.Exit(2)
	}
	arr := makeUlamArray(spiralDim, spiralStart)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%02d ", arr[i])
	}

}
