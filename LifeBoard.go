package main

import (
	"fmt"
	"math/rand"
	"os"
)

// LifeBoard - Grid for game of life slice of bool slices
type LifeBoard [][]bool // A slice of bool slices

// NewLifeBoard factory
func NewLifeBoard(numRows int, numCols int) *LifeBoard {
	board := make(LifeBoard, numRows)
	for i := 0; i < numRows; i++ {
		board[i] = make([]bool, numCols)
	}
	return &board
}

// Run i number of iterations
func (board LifeBoard) Run(i int) {
	// initialize/randomize board
	board.Randomize()
	for x := 0; x <= i; x++ {
		board.Update()
	}
}

// Randomize each cell of a life board
func (board LifeBoard) Randomize() {
	// for each set to random zero or 1
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			board[y][x] = (rand.Intn(2) != 0)
		}
	}
}

// Update board one cycle
func (board LifeBoard) Update() {
	// create new board to represent new state
	// for each cell
	// count neighbors

	//

	// write new state to new board

	// Copy new state to original board and discard new board
	// *** OR ***
	// Use new board, drop old one

}

// GetNeighborCount - How many live cells are surrounding the given coordinates
func (board LifeBoard) GetNeighborCount(y int, x int) int {
	//maxRow =
	//maxCol
	neighborCount := 0

	return neighborCount
}

// Save board to file
func (board LifeBoard) Save(filename string) {
	// Load the file
	myfile, err := os.OpenFile(filename, os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) { // File doesn't exist to open
			myfile, err = os.Create(filename)
			if err != nil {
				fmt.Println("Error creating new file: ", err)
			}
		}
	} else {
		fmt.Printf("Warning: File %s exists already. Overwriting.\n", filename)
	}
	defer myfile.Close()

	// Write the grid to file
	for y := 0; y < numGridRows; y++ {
		for x := 0; x < numGridCols; x++ {
			if board[y][x] == true {
				myfile.WriteString("0")
			} else {
				myfile.WriteString("1")
			}
		}
		myfile.WriteString("\n")
	}
}

// LoadFromFile
