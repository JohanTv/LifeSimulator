package main

import (
	"fmt"
)

const gridSize = 25

// John Conway's game of life simulator
// Will run many many iterations on a large board
// and then analyze to see if there are any stable structures
// that formed. Hopefully something interesting will!

// LifeBoard Grid for game of life
type LifeBoard struct {
	// 2d array
	//grid int[][]
	grid [gridSize][gridSize]bool
}

// print board to console

// write board to file

// clear board

// initialize board

// update board

// search board for stable structures

func main() {
	fmt.Println("Welcome")
	board := LifeBoard{}
	fmt.Printf("%v", board)
}
