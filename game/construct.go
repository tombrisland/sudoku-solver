package game

import (
	"errors"
	"fmt"
)

// FromSize creates a grid from a size component
func FromSize(size int) (*Game, error) {
	if size%GridSize != 0 {
		return nil, errors.New(fmt.Sprintf("Size must be divisible by %d", GridSize))
	}

	cells := make([][]int, size)
	options := make([][][]int, size)

	for x := 0; x < size; x++ {
		cells[x] = make([]int, size)
		options[x] = make([][]int, size)
	}

	return &Game{
		cells:   &cells,
		options: &options,
		size:    size,
	}, nil
}

func CopyGame(g *Game) *Game {
	// Copy the cells matrix
	cells := make([][]int, len(*g.cells))
	for i := range *g.cells {
		cells[i] = make([]int, len((*g.cells)[i]))
		copy(cells[i], (*g.cells)[i])
	}

	return &Game{
		cells: &cells,
		// Options matrix is lazily initialised
		options: nil,
		size:    g.size,
		valid:   false,
	}
}

// FromMatrix creates a grid from a matrix already defined
func FromMatrix(matrix [][]int) *Game {
	// Assume square matrix
	size := len(matrix)

	// Create blank options matrix
	options := make([][][]int, size)

	for x := 0; x < size; x++ {
		options[x] = make([][]int, size)
	}

	return &Game{
		cells:   &matrix,
		options: &options,
		size:    size,
	}
}
