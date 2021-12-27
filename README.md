## Sudoku solver

Recursive solver which operates on a `9x9` matrix. 

Plan to extend with photo upload and OCR to parse the puzzle.

### Run

There's a test and benchmark in `tests/solve_test.go`, run with `go test -bench=. ./tests`

### Improvements
* Only update affected options

Currently, all options are re-generated on setting a new value. Instead, could re-generate only the affected options.
* Make copying more efficient / copy less?

Every time a new value is attempted the matrix must be copied. Is there a way of compartmentalising to cut down the amount copied  

* Calculate game status alongside options

Calculating the game status requires an additional scan of the matrix which could be done during option generation.

* Add a better `String()` method

Include iterations to solve + other metrics