package main

import "fmt"

const numRuns = 10000
const numGridRows = 1000
const numGridCols = 1000

func main() {
	fmt.Println("Initializing Game of Life")
	board := NewLifeBoard(numGridRows, numGridCols)

	fmt.Println("Beginning simulation...")
	board.Run(numRuns)
	fmt.Println("Simulation complete...")

	board.Save("lifegrid.dat")
}
