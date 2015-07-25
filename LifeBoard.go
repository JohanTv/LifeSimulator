package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// LifeBoard - Grid for game of life slice of bool slices
type LifeBoard struct {
	grid             [][]uint8
	numCols, numRows int
}

// NewLifeBoard factory
func NewLifeBoard(numRows int, numCols int) *LifeBoard {
	board := LifeBoard{}
	board.numCols = numCols
	board.numRows = numRows
	board.grid = make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		board.grid[i] = make([]uint8, numCols)
	}
	return &board
}

// Run i number of iterations
func (board LifeBoard) Run(i int) {
	board.BigBang()
	for x := 0; x <= i; x++ {
		board.Update()
	}
}

// BigBang randomizes each cell of the life board
func (board LifeBoard) BigBang() {
	// for each set to random zero or 1
	for y := 0; y < len(board.grid); y++ {
		for x := 0; x < len(board.grid[0]); x++ {
			board.grid[y][x] = uint8(rand.Intn(2))
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

	var neighborCount uint8
	neighborCount = 0

	if x > 0 && y > 0 { // Top left
		neighborCount += board.grid[y-1][x-1]
	}
	if x > 0 { // Left
		neighborCount += board.grid[y][x-1]
	}
	if x > 0 && y < (board.numRows-1) { // Bottom left
		neighborCount += board.grid[y+1][x-1]
	}
	if y < (board.numRows - 1) { // Bottom
		neighborCount += board.grid[y+1][x]
	}
	if x < (board.numCols-1) && y < (board.numRows-1) { // Bottom Right
		neighborCount += board.grid[y+1][x+1]
	}
	if x < (board.numCols - 1) { // Right
		neighborCount += board.grid[y][x+1]
	}
	if x < (board.numCols-1) && y > 0 { // Top right
		neighborCount += board.grid[y-1][x+1]
	}

	return int(neighborCount)
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
			myfile.WriteString(strconv.Itoa(int(board.grid[y][x])))
		}
		myfile.WriteString("\n")
	}
}

// LoadFromFile
