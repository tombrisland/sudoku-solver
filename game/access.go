package game

// ForEachCell calls a consumer function for each cell
func (g *Game) ForEachCell(consumer func(int, int, int)) {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			consumer(x, y, g.Get(x, y))
		}
	}
}

// Set the value of a cell
func (g *Game) Set(x int, y int, value int) {
	// Set the cell with the value
	(*g.cells)[x][y] = value

	// Remove any options for the cell
	if g.options != nil {
		(*g.options)[x][y] = nil
	}

	// FindOptions are no longer valid
	g.valid = false
}

// Get returns the value of a cell
func (g *Game) Get(x int, y int) int {
	// Return the value at that cell position
	return (*g.cells)[x][y]
}

// GetOptions returns the possible values for a cell
func (g *Game) GetOptions(x int, y int) []int {
	if !g.valid {
		// Update valid options
		g.FindOptions()
	}

	// Return the value at that cell position
	return (*g.options)[x][y]
}

// GetGrid returns the upper left-hand coordinate of the parent grid
func (g *Game) GetGrid(x int, y int) (int, int) {
	// Calculate grid location via floor div
	return (x / GridSize) * GridSize, (y / GridSize) * GridSize
}

// Status of the game
type Status int

const (
	Solved   = 0
	Unsolved = 1
	Errored  = 2
)

// Status returns true when the game is Solved successfully
func (g *Game) Status() Status {
	// Total options available
	zeroed := 0
	options := 0

	g.ForEachCell(func(x int, y int, val int) {
		if val == 0 {
			zeroed += 1
		}

		options += len(g.GetOptions(x, y))
	})

	if options == 0 {
		// Solved when 0 options and no zeroed cells
		if zeroed == 0 {
			return Solved
		}
		return Errored
	}
	return Unsolved
}
