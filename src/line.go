package src

func (b *Board) Row(number int) []*Cell {
	cells := make([]*Cell, 0)
	for _, cell := range b.Cells {
		if cell.Y == number {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (b *Board) Column(number int) []*Cell {
	cells := make([]*Cell, 0)
	for _, cell := range b.Cells {
		if cell.X == number {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (b *Board) Line(number int, direction string) []*Cell {
	if direction == "x" || direction == "row" {
		return b.Row(number)
	} else {
		return b.Column(number)
	}
	return nil
}

func (b *Board) rowContainsValue(number int, value int) bool {
	return cellSliceContainsValue(b.Row(number), value)
}

func (b *Board) columnContainsValue(number int, value int) bool {
	return cellSliceContainsValue(b.Column(number), value)
}
