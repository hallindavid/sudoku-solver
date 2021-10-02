package src

import "fmt"

type Board struct {
	Cells []*Cell
}

func BuildBoard() Board {
	fmt.Print("Creating Board......")
	b := Board{
		Cells: make([]*Cell, 0),
	}
	fmt.Print("done\n")

	// Create the Cells
	fmt.Print("Creating default board cells......")
	b.createBoardCells()
	fmt.Printf("done. %d cells created\n", len(b.Cells))

	// Fill with our basic sample values
	fmt.Print("Filling Board with sample values......")
	b.fillBoardWithSampleValues()
	fmt.Printf("done. %d cells pre-solved\n", b.SolvedCount())

	return b
}

func (b *Board) SolvedCount() int {

	var solved int = 0

	for _, c := range b.Cells {
		if c.Filled {
			solved++
		}
	}
	return solved
}

func (b *Board) PreSolvedCount() int {

	var solved int = 0

	for _, c := range b.Cells {
		if c.Filled && c.Default {
			solved++
		}
	}
	return solved
}

func (b *Board) UnsolvedCount() int {
	return 81 - b.SolvedCount()
}
