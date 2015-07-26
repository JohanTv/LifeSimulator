package main

import "fmt"

const filename string = "lifegrid.dat"
const numRuns int = 10000
const numGridRows int = 1000
const numGridCols int = 1000

func main() {
	fmt.Println("Initializing Game of Life")

	board := NewLifeBoard(numGridRows, numGridCols)
	board.Run(numRuns)

	fmt.Println("Simulation complete... Writing to file.")

	board.Save(filename)

	fmt.Println("Data written to " + filename)
}
