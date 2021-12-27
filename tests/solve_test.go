package tests

import (
	"sudoku/game"
	"testing"
)

// Return a new instance of the test matrix
func matrix() [][]int {
	return [][]int{
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 4, 9, 2, 0, 1, 3, 6, 0},
		{0, 8, 0, 4, 0, 6, 0, 1, 0},
		{0, 1, 4, 0, 0, 0, 2, 7, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 7, 2, 0, 0, 0, 1, 5, 0},
		{0, 2, 0, 1, 0, 9, 0, 4, 0},
		{0, 9, 0, 5, 0, 8, 6, 3, 0},
		{0, 0, 0, 0, 3, 0, 0, 0, 0},
	}
}

// Ensure we can solve a valid puzzle
func TestSolveSuccess(t *testing.T) {
	g := game.FromMatrix(matrix())

	g.Solve()

	println(g.String())

	if g.Status() != game.Solved {
		t.FailNow()
	}
}

// Test performance of Solve
func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Initialise the game
		b.StopTimer()
		g := game.FromMatrix(matrix())
		b.StartTimer()

		// Time only the solve function
		g.Solve()
	}
}
