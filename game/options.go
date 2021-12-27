package game

import (
	"sync"
	"sync/atomic"
)

// Allowed values in sudoku
var allowed = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// The onlyOption available in a slice
const onlyOption = 0

// FindOptions updates the valid values for each cell
func (g *Game) FindOptions() {
	var wg sync.WaitGroup

	// Create the option matrix if it doesn't exist
	if g.options == nil {
		o := make([][][]int, g.size)

		g.options = &o
	}

	// How many options in total
	var opts uint64

	// Generate options for each cell
	g.ForEachCell(func(aX int, aY int, aCell int) {
		wg.Add(1)

		// Ensure an options array exists for each row
		if (*g.options)[aX] == nil {
			(*g.options)[aX] = make([][]int, g.size)
		}

		// Goroutine for each cell in matrix
		go func(wg *sync.WaitGroup, aX int, aY int, aCell int) {
			defer wg.Done()

			// Only need to find empty cells
			if aCell != 0 {
				return
			}

			avail := optionsForCell(g, aX, aY)

			// Set the cell options
			(*g.options)[aX][aY] = avail

			// Increment the options count
			atomic.AddUint64(&opts, uint64(len(avail)))
		}(&wg, aX, aY, aCell)
	})

	wg.Wait()

	// Options have been generated
	g.valid = true
}

func optionsForCell(g *Game, aX int, aY int) []int {
	// Find the grid for the cell
	aGridX, aGridY := g.GetGrid(aX, aY)

	// Numbers used already in col, row or grid
	used := make([]bool, 10)

	g.ForEachCell(func(bX int, bY int, bCell int) {
		bGridX, bGridY := g.GetGrid(bX, bY)

		// If col, row or grid matches
		if aX == bX || aY == bY || (aGridX == bGridX && aGridY == bGridY) {
			// Value has already been used
			used[bCell] = true
		}
	})

	avail := make([]int, 0)

	// Any unused values are available options for our cell
	for _, val := range allowed {
		if used[val] == false {
			avail = append(avail, val)
		}
	}

	return avail
}
