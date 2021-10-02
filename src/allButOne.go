package src

import "fmt"

// The ALL BUT ONE rule looks at a sector, row or column and if there are of these segments with 8/9 values, it will fill in the blank value with the only value left

func (b *Board) checkColumnsForAllButOne() string {
	for x := 1; x < 10; x++ {
		retVal := checkAllButOneCells(b.Column(x))

		if retVal != "" {
			return fmt.Sprintf("Column %d: All-but-one rule. %s", x, retVal)
		}
	}
	return ""
}

func (b *Board) checkRowsForAllButOne() string {
	for y := 1; y < 10; y++ {
		retVal := checkAllButOneCells(b.Row(y))

		if retVal != "" {
			return fmt.Sprintf("Row %d: All-but-one rule. %s", y, retVal)
		}
	}
	return ""
}

func (b *Board) checkSectorsForAllButOne() string {

	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			retVal := checkAllButOneCells(b.Sector(x, y))
			if retVal != "" {
				return fmt.Sprintf("Sector [%d, %d]: All-but-one rule. %s", x, y, retVal)
			}
		}
	}

	return ""
}

func checkAllButOneCells(cells []*Cell) string {
	var filledCount int = 0
	remainingValues := getValSlice()

	for _, c := range cells {
		if c.Filled {
			filledCount++

			for i, v := range remainingValues {
				if v == c.Value {
					remainingValues[i] = remainingValues[len(remainingValues)-1] // Copy last element to index i.
					remainingValues[len(remainingValues)-1] = 0                  // Erase last element (write zero value).
					remainingValues = remainingValues[:len(remainingValues)-1]   // Truncate slice
				}
			}
		}
	}

	if filledCount == 8 {
		// Find the Unfilled value
		for _, c := range cells {
			if !c.Filled {
				c.Value = remainingValues[0] // Fill with remaining element in values array
				return fmt.Sprintf("[%d,%d] = %d", c.X, c.Y, c.Value)
			}
		}
	}
	return ""
}

func getValSlice() []int {
	values := make([]int, 9)
	for i := 0; i < 9; i++ {
		values[i] = i + 1
	}
	return values
}
