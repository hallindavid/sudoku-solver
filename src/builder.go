package src

func (b *Board) fillBoardWithSampleValues() {
	b.fillWithSample1()
}

func (b *Board) fillWithSample1() {
	b.fillCellAt(2, 1, 6, true)
	b.fillCellAt(4, 1, 9, true)
	b.fillCellAt(5, 1, 8, true)
	b.fillCellAt(6, 1, 3, true)
	b.fillCellAt(8, 1, 7, true)

	b.fillCellAt(2, 2, 9, true)
	b.fillCellAt(5, 2, 1, true)
	b.fillCellAt(7, 2, 6, true)
	b.fillCellAt(8, 2, 2, true)

	b.fillCellAt(3, 3, 4, true)
	b.fillCellAt(6, 3, 6, true)
	b.fillCellAt(9, 3, 3, true)

	b.fillCellAt(1, 4, 2, true)

	b.fillCellAt(2, 6, 4, true)
	b.fillCellAt(7, 6, 7, true)

	b.fillCellAt(4, 7, 2, true)
	b.fillCellAt(6, 7, 1, true)
	b.fillCellAt(7, 7, 3, true)

	b.fillCellAt(1, 8, 1, true)
	b.fillCellAt(2, 8, 7, true)
	b.fillCellAt(6, 8, 8, true)
	b.fillCellAt(8, 8, 9, true)
	b.fillCellAt(9, 8, 2, true)

	b.fillCellAt(5, 9, 3, true)
	b.fillCellAt(7, 9, 4, true)
	b.fillCellAt(8, 9, 6, true)
}
