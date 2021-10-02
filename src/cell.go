package src

import (
	"errors"
	"fmt"
)

type Cell struct {
	X       int
	Y       int
	Filled  bool
	Value   int
	Default bool
}

func (b *Board) createBoardCells() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			b.Cells = append(b.Cells, &Cell{
				X:       i,
				Y:       j,
				Value:   0,
				Filled:  false,
				Default: false,
			})
		}
	}
}

func (b *Board) Cell(x int, y int) *Cell {
	for _, c := range b.Cells {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

func (b *Board) getCellVal(x int, y int) string {
	for _, c := range b.Cells {
		if c.X == x && c.Y == y {
			return c.Val()
		}
	}
	return ""
}

func (b *Board) fillCellAt(x int, y int, value int, isDefault bool) error {
	if value < 1 || value > 9 {
		return errors.New("Value must be between 1 and 9")
	}

	for _, c := range b.Cells {
		if c.X == x && c.Y == y {
			if c.Filled {
				return errors.New(fmt.Sprintf("Cell [%d, %d] already has a value.", x, y))
			}
			c.Value = value
			c.Filled = true
			c.Default = isDefault
			return nil
		}
	}
	return errors.New("Unable to find cell at that location")
}

func (c *Cell) fill(value int) error {
	if value < 1 || value > 9 {
		return errors.New("Value must be between 1 and 9")
	}

	if c.Filled {
		return errors.New(fmt.Sprintf("Cell [%d, %d] already has a value", c.X, c.Y))
	}

	c.Value = value
	c.Filled = true
	c.Default = false

	return nil
}

func (c *Cell) Val() string {
	if c.Filled {
		return fmt.Sprintf("%d", c.Value)
	}
	return " "
}

// Cell in Line/Sector Getters and Checkers
func cellSliceContainsValue(cells []*Cell, value int) bool {
	for _, c := range cells {
		if c.Value == value {
			return true
		}
	}
	return false
}

func cellInSliceWithValue(cells []*Cell, value int) *Cell {
	for _, c := range cells {
		if c.Value == value {
			return c
		}
	}
	return nil
}

func (b *Board) cellInRowWithValue(number int, value int) *Cell {
	return cellInSliceWithValue(b.Row(number), value)
}

func (b *Board) cellInColumnWithValue(number int, value int) *Cell {
	return cellInSliceWithValue(b.Column(number), value)
}

func (b *Board) cellInSectorWithValue(x int, y int, value int) *Cell {
	return cellInSliceWithValue(b.Sector(x, y), value)
}

func solvedCells(cells []*Cell) []*Cell {
	filled := make([]*Cell, 0)
	for _, c := range cells {
		if c.Filled {
			filled = append(filled, c)
		}
	}
	return filled
}

func (c *Cell) getSectorXY() (x int, y int) {
	x = 1
	if c.X > 3 {
		x = 2
	}
	if c.X > 6 {
		x = 3
	}

	y = 1
	if c.Y > 3 {
		y = 2
	}
	if c.Y > 6 {
		y = 3
	}
	return x, y
}

func (c *Cell) getSectorX() (x int) {
	x, _ = c.getSectorXY()
	return x
}
func (c *Cell) getSectorY() (y int) {
	_, y = c.getSectorXY()
	return y
}
