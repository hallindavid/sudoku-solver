package solver

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintBoard(b Board) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("\033[2J")
	fmt.Println("*-------*-------*-------*")

	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.A1), suprint(b.A2), suprint(b.A3), suprint(b.A4), suprint(b.A5), suprint(b.A6), suprint(b.A7), suprint(b.A8), suprint(b.A9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.B1), suprint(b.B2), suprint(b.B3), suprint(b.B4), suprint(b.B5), suprint(b.B6), suprint(b.B7), suprint(b.B8), suprint(b.B9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.C1), suprint(b.C2), suprint(b.C3), suprint(b.C4), suprint(b.C5), suprint(b.C6), suprint(b.C7), suprint(b.C8), suprint(b.C9))
	fmt.Println("*-------*-------*-------*")
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.D1), suprint(b.D2), suprint(b.D3), suprint(b.D4), suprint(b.D5), suprint(b.D6), suprint(b.D7), suprint(b.D8), suprint(b.D9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.E1), suprint(b.E2), suprint(b.E3), suprint(b.E4), suprint(b.E5), suprint(b.E6), suprint(b.E7), suprint(b.E8), suprint(b.E9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.F1), suprint(b.F2), suprint(b.F3), suprint(b.F4), suprint(b.F5), suprint(b.F6), suprint(b.F7), suprint(b.F8), suprint(b.F9))
	fmt.Println("*-------*-------*-------*")
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.G1), suprint(b.G2), suprint(b.G3), suprint(b.G4), suprint(b.G5), suprint(b.G6), suprint(b.G7), suprint(b.G8), suprint(b.G9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.H1), suprint(b.H2), suprint(b.H3), suprint(b.H4), suprint(b.H5), suprint(b.H6), suprint(b.H7), suprint(b.H8), suprint(b.H9))
	fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n", suprint(b.I1), suprint(b.I2), suprint(b.I3), suprint(b.I4), suprint(b.I5), suprint(b.I6), suprint(b.I7), suprint(b.I8), suprint(b.I9))
	fmt.Println("*-------*-------*-------*")

	fmt.Printf("Solved Spots: %d / 81 \n\n", b.SolvedCount())
}

func suprint(i int) string {
	if i > 0 {
		return fmt.Sprintf("%d", i)
	}
	return " "
}
