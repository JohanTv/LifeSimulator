package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// LifeBoard - Grid for game of life slice of bool slices
type LifeBoard struct {
	gridA, gridB             [][]uint8
	activeGrid, inactiveGrid *[][]uint8

	numCols, numRows int
}

// NewLifeBoard factory
func NewLifeBoard(numRows int, numCols int) *LifeBoard {
	board := LifeBoard{}
	board.numCols = numCols
	board.numRows = numRows
	board.gridA = make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		board.gridA[i] = make([]uint8, numCols)
	}
	board.gridB = make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		board.gridB[i] = make([]uint8, numCols)
	}
	board.activeGrid = &board.gridA
	board.inactiveGrid = &board.gridB
	return &board
}

// Run i number of iterations
func (board *LifeBoard) Run(i int) {
	board.BigBang()
	for x := 0; x <= i; x++ {
		board.Update()
	}
}

// BigBang randomizes each cell of the life board
func (board LifeBoard) BigBang() {
	rand.Seed(time.Now().UTC().UnixNano())
	// for each set to random zero or 1
	for y := 0; y < board.numRows; y++ {
		for x := 0; x < board.numCols; x++ {
			(*board.activeGrid)[y][x] = uint8(rand.Intn(2))
		}
	}
}

// Update board one cycle
func (board *LifeBoard) Update() {

	var neighborCount int

	// Determine fate of cell
	for y := 0; y < board.numRows; y++ {
		for x := 0; x < board.numCols; x++ {
			neighborCount = board.GetNeighborCount(y, x)
			if (*board.activeGrid)[y][x] == 1 { // Alive
				if neighborCount < 2 {
					(*board.inactiveGrid)[y][x] = 0
				} else if neighborCount == 2 || neighborCount == 3 {
					(*board.inactiveGrid)[y][x] = 1
				} else if neighborCount > 3 {
					(*board.inactiveGrid)[y][x] = 0
				}
			} else { // Dead
				if neighborCount == 3 {
					(*board.inactiveGrid)[y][x] = 1
				} else {
					(*board.inactiveGrid)[y][x] = 0
				}
			}
		}
	}

	// Swap active and inactive grids
	if board.activeGrid == &board.gridA {
		board.activeGrid = &board.gridB
		board.inactiveGrid = &board.gridA
	} else {
		board.activeGrid = &board.gridA
		board.inactiveGrid = &board.gridB
	}

}

// GetNeighborCount - How many live cells are surrounding the given coordinates
func (board LifeBoard) GetNeighborCount(y int, x int) int {

	var neighborCount uint8

	if x > 0 && y > 0 { // Top left
		neighborCount += (*board.activeGrid)[y-1][x-1]
	}
	if x > 0 { // Left
		neighborCount += (*board.activeGrid)[y][x-1]
	}
	if x > 0 && y < (board.numRows-1) { // Bottom left
		neighborCount += (*board.activeGrid)[y+1][x-1]
	}
	if y < (board.numRows - 1) { // Bottom
		neighborCount += (*board.activeGrid)[y+1][x]
	}
	if x < (board.numCols-1) && y < (board.numRows-1) { // Bottom Right
		neighborCount += (*board.activeGrid)[y+1][x+1]
	}
	if x < (board.numCols - 1) { // Right
		neighborCount += (*board.activeGrid)[y][x+1]
	}
	if x < (board.numCols-1) && y > 0 { // Top right
		neighborCount += (*board.activeGrid)[y-1][x+1]
	}

	return int(neighborCount)
}

// Save board to file
func (board LifeBoard) Save(filename string) {
	// Load the file
	myfile, err := os.OpenFile(filename, os.O_CREATE, 0755)
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
			myfile.WriteString(strconv.Itoa(int((*board.activeGrid)[y][x])))
		}
		myfile.WriteString("\r\n")
	}
}

// LoadFromFile
