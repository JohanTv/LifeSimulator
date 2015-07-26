package main

import "fmt"

const filename string = "lifegrid.dat"

// @TODO Accept these are command line arguments
const numRuns int = 1000
const numGridRows int = 224
const numGridCols int = 222

func main() {
	fmt.Println("Initializing and running Game of Life")

	board := NewLifeBoard(numGridRows, numGridCols)
	board.Run(numRuns)

	fmt.Println("Simulation complete... Writing to file.")

	board.Save(filename)

	fmt.Println("Data written to " + filename)
}
