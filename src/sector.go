package src

// sectors X and Y properties are in the range of 1-3 (9 sectors in a sudoku board)
func (b *Board) Sector(x int, y int) []*Cell {
	xVals := sectorToSlice(x)
	yVals := sectorToSlice(y)
	cells := make([]*Cell, 0)

	for _, cell := range b.Cells {
		acceptableX, acceptableY := false, false

		for _, acceptedXRange := range xVals {
			if acceptedXRange == cell.X {
				acceptableX = true
			}
		}

		for _, acceptedYRange := range yVals {
			if acceptedYRange == cell.Y {
				acceptableY = true
			}
		}

		if acceptableY && acceptableX {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (b *Board) sectorAtXYContainsValue(x int, y int, value int) bool {
	return cellSliceContainsValue(b.Sector(x, y), value)
}

func (b *Board) horizontallyRelatedSectors(sectorY int) [][]*Cell {
	sectors := make([][]*Cell, 0)

	sectors = append(sectors, b.Sector(1, sectorY))
	sectors = append(sectors, b.Sector(2, sectorY))
	sectors = append(sectors, b.Sector(3, sectorY))
	return sectors
}

func (b *Board) verticallyRelatedSectors(sectorX int) [][]*Cell {
	sectors := make([][]*Cell, 0)

	sectors = append(sectors, b.Sector(sectorX, 1))
	sectors = append(sectors, b.Sector(sectorX, 2))
	sectors = append(sectors, b.Sector(sectorX, 3))
	return sectors
}

func sectorToSlice(input int) []int {
	if input == 1 {
		return []int{1, 2, 3}
	}
	if input == 2 {
		return []int{4, 5, 6}
	}
	if input == 3 {
		return []int{7, 8, 9}
	}

	return make([]int, 0)
}
