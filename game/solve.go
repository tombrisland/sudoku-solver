package game

// findMostLikely returns the cell with the lowest number of options
func (g *Game) findMostLikely() (int, int) {
	// 1 means cell value is known
	lowest, lowX, lowY := 9, -1, -1

	g.ForEachCell(func(x int, y int, cell int) {
		options := g.GetOptions(x, y)
		count := len(options)

		// Find first single value or lowest
		if count > 0 && count < lowest {
			lowest, lowX, lowY = count, x, y
		}
	})

	return lowX, lowY
}

// Solve the game and update the matrix cells
func (g *Game) Solve() {
	solved := solveRecursively(g)
	g.cells = solved.cells
	g.options = solved.options
}

// Recurse the available options and solve the game
func solveRecursively(g *Game) *Game {
	switch g.Status() {
	case Solved:
		return g
	case Unsolved:
		x, y := g.findMostLikely()
		options := g.GetOptions(x, y)

		// If only a single option set it
		if len(options) == 1 {
			c := CopyGame(g)
			c.Set(x, y, options[onlyOption])

			return solveRecursively(c)
		}

		// Otherwise, recursively attempt options
		for _, option := range options {
			// Copy the game to avoid overwriting
			c := CopyGame(g)
			c.Set(x, y, option)

			// Try to solve with the option
			c = solveRecursively(c)

			// Only return successful branches
			if c != nil {
				return c
			}
		}
	}

	return nil
}
