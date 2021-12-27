package game

import (
	"strconv"
)

const GridSize = 3

// Game of sudoku
type Game struct {
	// The game made up of cells
	cells *[][]int

	// The size of the square cell matrix
	size int

	// Current options for each cell
	options *[][][]int

	// Whether the options structure is valid
	valid bool
}

// String returns a printable representation of the game
func (g *Game) String() string {
	var s string

	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			s += strconv.Itoa(g.Get(x, y)) + " "
		}
		s += "\n"
	}

	return s
}
