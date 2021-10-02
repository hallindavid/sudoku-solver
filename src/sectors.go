package solver

import (
	"fmt"
	"reflect"
)

func (b *Board) lineContains(val int, row int, direction string) (bool, string) {
	var line = make([]*int, 9)
	line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8] = b.Line(row, direction)

	for _, value := range line {
		if *value == val {
			return true, "A1"
		}
	}
	return false, ""
}

func (b *Board) sectorContains(val int, x int, y int) (bool, string) {
	var sector = make([]*int, 9)
	sector[0], sector[1], sector[2], sector[3], sector[4], sector[5], sector[6], sector[7], sector[8] = b.Sector(x, y)

	for _, value := range sector {
		if *value == val {
			return true, "A1"
		}
	}
	return false, ""
}

func (b *Board) getField(field string) int {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func (b *Board) Cell(x int, y int) int {
	var propName string

	switch x {
	case 1:
		propName = fmt.Sprintf("%s%d", "A", y)
	case 2:
		propName = fmt.Sprintf("%s%d", "B", y)
	case 3:
		propName = fmt.Sprintf("%s%d", "C", y)
	case 4:
		propName = fmt.Sprintf("%s%d", "D", y)
	case 5:
		propName = fmt.Sprintf("%s%d", "E", y)
	case 6:
		propName = fmt.Sprintf("%s%d", "F", y)
	case 7:
		propName = fmt.Sprintf("%s%d", "G", y)
	case 8:
		propName = fmt.Sprintf("%s%d", "H", y)
	case 9:
		propName = fmt.Sprintf("%s%d", "I", y)
	}

	return b.getField(propName)

}

func (b *Board) Line(r int, direction string) (*int, *int, *int, *int, *int, *int, *int, *int, *int) {
	if direction == "x" {
		if r == 1 {
			return &b.A1, &b.A2, &b.A3, &b.A4, &b.A5, &b.A6, &b.A7, &b.A8, &b.A9
		}
		if r == 2 {
			return &b.B1, &b.B2, &b.B3, &b.B4, &b.B5, &b.B6, &b.B7, &b.B8, &b.B9
		}
		if r == 3 {
			return &b.C1, &b.C2, &b.C3, &b.C4, &b.C5, &b.C6, &b.C7, &b.C8, &b.C9
		}
		if r == 4 {
			return &b.D1, &b.D2, &b.D3, &b.D4, &b.D5, &b.D6, &b.D7, &b.D8, &b.D9
		}
		if r == 5 {
			return &b.E1, &b.E2, &b.E3, &b.E4, &b.E5, &b.E6, &b.E7, &b.E8, &b.E9
		}
		if r == 6 {
			return &b.F1, &b.F2, &b.F3, &b.F4, &b.F5, &b.F6, &b.F7, &b.F8, &b.F9
		}

		if r == 7 {
			return &b.G1, &b.G2, &b.G3, &b.G4, &b.G5, &b.G6, &b.G7, &b.G8, &b.G9
		}
		if r == 8 {
			return &b.H1, &b.H2, &b.H3, &b.H4, &b.H5, &b.H6, &b.H7, &b.H8, &b.H9
		}
		if r == 9 {
			return &b.I1, &b.I2, &b.I3, &b.I4, &b.I5, &b.I6, &b.I7, &b.I8, &b.I9
		}
	} else if direction == "y" {
		if r == 1 {
			return &b.A1, &b.B1, &b.C1, &b.D1, &b.E1, &b.F1, &b.G1, &b.H1, &b.I1
		}
		if r == 2 {
			return &b.A2, &b.B2, &b.C2, &b.D2, &b.E2, &b.F2, &b.G2, &b.H2, &b.I2
		}
		if r == 3 {
			return &b.A3, &b.B3, &b.C3, &b.D3, &b.E3, &b.F3, &b.G3, &b.H3, &b.I3
		}

		if r == 4 {
			return &b.A4, &b.B4, &b.C4, &b.D4, &b.E4, &b.F4, &b.G4, &b.H4, &b.I4
		}
		if r == 5 {
			return &b.A5, &b.B5, &b.C5, &b.D5, &b.E5, &b.F5, &b.G5, &b.H5, &b.I5
		}
		if r == 6 {
			return &b.A6, &b.B6, &b.C6, &b.D6, &b.E6, &b.F6, &b.G6, &b.H6, &b.I6
		}

		if r == 7 {
			return &b.A7, &b.B7, &b.C7, &b.D7, &b.E7, &b.F7, &b.G7, &b.H7, &b.I7
		}
		if r == 8 {
			return &b.A8, &b.B8, &b.C8, &b.D8, &b.E8, &b.F8, &b.G8, &b.H8, &b.I8
		}
		if r == 9 {
			return &b.A9, &b.B9, &b.C9, &b.D9, &b.E9, &b.F9, &b.G9, &b.H9, &b.I9
		}
	}
}

func (b *Board) Sector(x int, y int) (*int, *int, *int, *int, *int, *int, *int, *int, *int) {
	if x == 1 && y == 1 {
		return &b.A1, &b.A2, &b.A3, &b.B1, &b.B2, &b.B3, &b.C1, &b.C2, &b.C3
	}
	if x == 1 && y == 2 {
		return &b.A4, &b.A5, &b.A6, &b.B4, &b.B5, &b.B6, &b.C4, &b.C5, &b.C6
	}
	if x == 1 && y == 3 {
		return &b.A7, &b.A8, &b.A9, &b.B7, &b.B8, &b.B9, &b.C7, &b.C8, &b.C9
	}

	if x == 2 && y == 1 {
		return &b.D1, &b.D2, &b.D3, &b.E1, &b.E2, &b.E3, &b.F1, &b.F2, &b.F3
	}
	if x == 2 && y == 2 {
		return &b.D4, &b.D5, &b.D6, &b.E4, &b.E5, &b.E6, &b.F4, &b.F5, &b.F6
	}
	if x == 2 && y == 3 {
		return &b.D7, &b.D8, &b.D9, &b.E7, &b.E8, &b.E9, &b.F7, &b.F8, &b.F9
	}

	if x == 3 && y == 1 {
		return &b.G1, &b.G2, &b.G3, &b.H1, &b.H2, &b.H3, &b.I1, &b.I2, &b.I3
	}
	if x == 3 && y == 2 {
		return &b.G4, &b.G5, &b.G6, &b.H4, &b.H5, &b.H6, &b.I4, &b.I5, &b.I6
	}
	if x == 3 && y == 3 {
		return &b.G7, &b.G8, &b.G9, &b.H7, &b.H8, &b.H9, &b.I7, &b.I8, &b.I9
	}
}
