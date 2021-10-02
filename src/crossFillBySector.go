package src

import "fmt"

func (b *Board) checkCrossFillBySector() string {
	for _, c := range solvedCells(b.Cells) {
		sectorX, sectorY := c.getSectorXY()
		//_, sectorY := c.getSectorXY()

		// Get horizontally related sectors
		horizontalSectors := b.horizontallyRelatedSectors(sectorY)
		verticalSectors := b.verticallyRelatedSectors(sectorX)

		if checkNumberOfSectorsWithValue(horizontalSectors, c.Value) == 2 {
			//fmt.Printf("\nPossible Horizontal Crossfill on cell [%d, %d]:%d\n", c.X, c.Y, c.Value)
			retHVal := b.attemptHorizontalCrossFill(horizontalSectors, c.Value)
			if retHVal != "" {
				//fmt.Printf("Successful Crossfill [%d, %d]:%d : %s \n", c.X, c.Y, c.Value, retHVal)
				return fmt.Sprintf("Horizontal Crossfill: %s", retHVal)
			}
		}

		if checkNumberOfSectorsWithValue(verticalSectors, c.Value) == 2 {
			//fmt.Printf("\nPossible Vertical Crossfill on cell [%d, %d]:%d\n", c.X, c.Y, c.Value)
			retVVal := b.attemptVerticalCrossFill(verticalSectors, c.Value)
			if retVVal != "" {
				//fmt.Printf("Successful Crossfill [%d, %d]:%d : %s \n", c.X, c.Y, c.Value, retVVal)
				return fmt.Sprintf("Vertical Crossfill: %s", retVVal)
			}
		}
	}
	return ""
}

func checkNumberOfSectorsWithValue(sectors [][]*Cell, value int) int {
	counter := 0
	for _, s := range sectors {
		if cellSliceContainsValue(s, value) {
			counter++
		}
	}
	return counter
}

func (b *Board) attemptHorizontalCrossFill(sectors [][]*Cell, value int) string {
	sectorRowSpan := sectorToSlice(sectors[0][0].getSectorY())

	var sectorWithoutValue []*Cell

	for _, s := range sectors {
		if cellSliceContainsValue(s, value) {
			c := cellInSliceWithValue(s, value)

			for index, num := range sectorRowSpan {
				if num == c.Y {
					sectorRowSpan[index] = sectorRowSpan[len(sectorRowSpan)-1] // Copy last element to index i.
					sectorRowSpan[len(sectorRowSpan)-1] = 0                    // Erase last element (write zero value).
					sectorRowSpan = sectorRowSpan[:len(sectorRowSpan)-1]       // Truncate slice
				}
			}
		} else {
			sectorWithoutValue = s
		}
	}

	//At this point, we should have a sector that's missing the value, as well as the row (sectorRowSpan[0]) where the value belongs
	// Now we need to see if there is only 1 cell that does not have data, and if that's the one, then we can fill it with the value
	possibleCells := make([]*Cell, 0)

	for _, c := range sectorWithoutValue {
		if c.Y == sectorRowSpan[0] && c.Filled != true { // it needs to be on the row where the other 2 sectors don't have data
			if !b.columnContainsValue(c.X, value) {
				possibleCells = append(possibleCells, c) // Add the cell if it isn't vertically excluded
			}
		}
	}

	if len(possibleCells) == 1 {
		possibleCells[0].fill(value)
		return fmt.Sprintf("[%d,%d] = %d", possibleCells[0].X, possibleCells[0].Y, possibleCells[0].Value)
	}

	return ""
}

func (b *Board) attemptVerticalCrossFill(sectors [][]*Cell, value int) string {

	sectorColSpan := sectorToSlice(sectors[0][0].getSectorX())

	var sectorWithoutValue []*Cell

	for _, s := range sectors {
		if cellSliceContainsValue(s, value) {
			c := cellInSliceWithValue(s, value)

			for index, num := range sectorColSpan {
				if num == c.X {
					sectorColSpan[index] = sectorColSpan[len(sectorColSpan)-1] // Copy last element to index i.
					sectorColSpan[len(sectorColSpan)-1] = 0                    // Erase last element (write zero value).
					sectorColSpan = sectorColSpan[:len(sectorColSpan)-1]       // Truncate slice
				}
			}
		} else {
			sectorWithoutValue = s
		}
	}

	//At this point, we should have a sector that's missing the value, as well as the column (sectorColSpan[0]) where the value belongs
	// Now we need to see if there is only 1 cell that does not have data, and if that's the one, then we can fill it with the value
	possibleCells := make([]*Cell, 0)

	for _, c := range sectorWithoutValue {
		if c.X == sectorColSpan[0] && c.Filled != true { // it needs to be on the column where the other 2 sectors don't have data
			if !b.rowContainsValue(c.Y, value) {
				possibleCells = append(possibleCells, c) // Add the cell to possibilities if it isn't horizontally excluded
			}
		}
	}

	if len(possibleCells) == 1 {
		possibleCells[0].fill(value)
		return fmt.Sprintf("[%d,%d] = %d", possibleCells[0].X, possibleCells[0].Y, possibleCells[0].Value)
	}

	return ""
}
