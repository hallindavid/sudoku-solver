package src

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func (b *Board) Print() {

	fmt.Print("\n\n")

	printGuideLine()

	for y := 1; y <= 9; y++ {
		//Row Start
		fmt.Print("|")

		for x := 1; x <= 9; x++ {
			c := b.Cell(x, y)
			if c.Default {
				fmt.Printf(" %s ", color.GreenString(c.Val()))
			} else {
				fmt.Printf(" %s ", c.Val())
			}

			if x == 3 || x == 6 || x == 9 {
				fmt.Print("|")
			}
		}

		// Row End
		fmt.Print("\n")

		// Print Sector Dividing Line
		if y == 3 || y == 6 || y == 9 {
			printGuideLine()
		}
	}

	fmt.Printf("%s | %d / 81 Solved\n\n", color.GreenString("%d pre-solved", b.PreSolvedCount()), b.SolvedCount())

	//b.PrintByRow()
	//b.PrintByCol()

}

func (b *Board) PrintByRow() {

	for i := 1; i < 10; i++ {
		fmt.Printf("Row %d Data \n", i)
		row := b.Row(i)
		for _, c := range row {
			fmt.Printf("[%d, %d]: %s |", c.X, c.Y, c.Val())
		}
		fmt.Printf("\n")
	}
}

func (b *Board) PrintByCol() {

	for i := 1; i < 10; i++ {
		fmt.Printf("Column %d Data \n", i)
		col := b.Column(i)
		for _, c := range col {
			fmt.Printf("[%d, %d]: %s |", c.X, c.Y, c.Val())
		}
		fmt.Printf("\n")
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("\033[2J")
}

func printGuideLine() {
	fmt.Println("*---------*---------*---------*")
}

func (b *Board) dumpCells() {
	for i, c := range b.Cells {
		fmt.Printf("%02d: [%d, %d]: %d \n", i, c.X, c.Y, c.Value)
	}
}
