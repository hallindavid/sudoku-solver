package solver

func (b *Board) checkLines() {
	//Horizontal
	finishGroup(b.Line(1, "x"))
	finishGroup(b.Line(2, "x"))
	finishGroup(b.Line(3, "x"))
	finishGroup(b.Line(4, "x"))
	finishGroup(b.Line(5, "x"))
	finishGroup(b.Line(6, "x"))
	finishGroup(b.Line(7, "x"))
	finishGroup(b.Line(8, "x"))
	finishGroup(b.Line(9, "x"))

	finishGroup(b.Line(1, "y"))
	finishGroup(b.Line(2, "y"))
	finishGroup(b.Line(3, "y"))
	finishGroup(b.Line(4, "y"))
	finishGroup(b.Line(5, "y"))
	finishGroup(b.Line(6, "y"))
	finishGroup(b.Line(7, "y"))
	finishGroup(b.Line(8, "y"))
	finishGroup(b.Line(9, "y"))

	//finishGroup(&b.A1, &b.A2, &b.A3, &b.A4, &b.A5, &b.A6, &b.A7, &b.A8, &b.A9)
	//finishGroup(&b.B1, &b.B2, &b.B3, &b.B4, &b.B5, &b.B6, &b.B7, &b.B8, &b.B9)
	//finishGroup(&b.C1, &b.C2, &b.C3, &b.C4, &b.C5, &b.C6, &b.C7, &b.C8, &b.C9)
	//
	//finishGroup(&b.D1, &b.D2, &b.D3, &b.D4, &b.D5, &b.D6, &b.D7, &b.D8, &b.D9)
	//finishGroup(&b.E1, &b.E2, &b.E3, &b.E4, &b.E5, &b.E6, &b.E7, &b.E8, &b.E9)
	//finishGroup(&b.F1, &b.F2, &b.F3, &b.F4, &b.F5, &b.F6, &b.F7, &b.F8, &b.F9)
	//
	//finishGroup(&b.G1, &b.G2, &b.G3, &b.G4, &b.G5, &b.G6, &b.G7, &b.G8, &b.G9)
	//finishGroup(&b.H1, &b.H2, &b.H3, &b.H4, &b.H5, &b.H6, &b.H7, &b.H8, &b.H9)
	//finishGroup(&b.I1, &b.I2, &b.I3, &b.I4, &b.I5, &b.I6, &b.I7, &b.I8, &b.I9)
	//
	////Vertical
	//finishGroup(&b.A1, &b.B1, &b.C1, &b.D1, &b.E1, &b.F1, &b.G1, &b.H1, &b.H1)
	//finishGroup(&b.A2, &b.B2, &b.C2, &b.D2, &b.E2, &b.F2, &b.G2, &b.H2, &b.H2)
	//finishGroup(&b.A3, &b.B3, &b.C3, &b.D3, &b.E3, &b.F3, &b.G3, &b.H3, &b.H3)
	//
	//finishGroup(&b.A4, &b.B4, &b.C4, &b.D4, &b.E4, &b.F4, &b.G4, &b.H4, &b.H4)
	//finishGroup(&b.A5, &b.B5, &b.C5, &b.D5, &b.E5, &b.F5, &b.G5, &b.H5, &b.H5)
	//finishGroup(&b.A6, &b.B6, &b.C6, &b.D6, &b.E6, &b.F6, &b.G6, &b.H6, &b.H6)
	//
	//finishGroup(&b.A7, &b.B7, &b.C7, &b.D7, &b.E7, &b.F7, &b.G7, &b.H7, &b.H7)
	//finishGroup(&b.A8, &b.B8, &b.C8, &b.D8, &b.E8, &b.F8, &b.G8, &b.H8, &b.H8)
	//finishGroup(&b.A9, &b.B9, &b.C9, &b.D9, &b.E9, &b.F9, &b.G9, &b.H9, &b.H9)
}

func (b *Board) checkSectors() {
	finishGroup(b.Sector(1,1))
	finishGroup(b.Sector(1,2))
	finishGroup(b.Sector(1,3))

	finishGroup(b.Sector(2,1))
	finishGroup(b.Sector(2,2))
	finishGroup(b.Sector(2,3))

	finishGroup(b.Sector(3,1))
	finishGroup(b.Sector(3,2))
	finishGroup(b.Sector(3,3))

	//finishGroup(&b.A1, &b.A2, &b.A3, &b.B1, &b.B2, &b.B3, &b.C1, &b.C2, &b.C3)
	//finishGroup(&b.D1, &b.D2, &b.D3, &b.E1, &b.E2, &b.E3, &b.F1, &b.F2, &b.F3)
	//finishGroup(&b.G1, &b.G2, &b.G3, &b.H1, &b.H2, &b.H3, &b.I1, &b.I2, &b.I3)
	//
	//finishGroup(&b.A4, &b.A5, &b.A6, &b.B4, &b.B5, &b.B6, &b.C4, &b.C5, &b.C6)
	//finishGroup(&b.D4, &b.D5, &b.D6, &b.E4, &b.E5, &b.E6, &b.F4, &b.F5, &b.F6)
	//finishGroup(&b.G4, &b.G5, &b.G6, &b.H4, &b.H5, &b.H6, &b.I4, &b.I5, &b.I6)
	//
	//finishGroup(&b.A7, &b.A8, &b.A9, &b.B7, &b.B8, &b.B9, &b.C7, &b.C8, &b.C9)
	//finishGroup(&b.D7, &b.D8, &b.D9, &b.E7, &b.E8, &b.E9, &b.F7, &b.F8, &b.F9)
	//finishGroup(&b.G7, &b.G8, &b.G9, &b.H7, &b.H8, &b.H9, &b.I7, &b.I8, &b.I9)
}

func finishGroup(a *int, b *int, c *int, d *int, e *int, f *int, g *int, h *int, i *int) {
	found1, found2, found3 := false, false, false
	found4, found5, found6 := false, false, false
	found7, found8, found9 := false, false, false

	foundCount := 0

	if *a == 1 || *b == 1 || *c == 1 || *d == 1 || *e == 1 || *f == 1 || *g == 1 || *h == 1 || *i == 1 {
		found1 = true
		foundCount++
	}

	if *a == 2 || *b == 2 || *c == 2 || *d == 2 || *e == 2 || *f == 2 || *g == 2 || *h == 2 || *i == 2 {
		found2 = true
		foundCount++
	}

	if *a == 3 || *b == 3 || *c == 3 || *d == 3 || *e == 3 || *f == 3 || *g == 3 || *h == 3 || *i == 3 {
		found3 = true
		foundCount++
	}

	if *a == 4 || *b == 4 || *c == 4 || *d == 4 || *e == 4 || *f == 4 || *g == 4 || *h == 4 || *i == 4 {
		found4 = true
		foundCount++
	}

	if *a == 5 || *b == 5 || *c == 5 || *d == 5 || *e == 5 || *f == 5 || *g == 5 || *h == 5 || *i == 5 {
		found5 = true
		foundCount++
	}

	if *a == 6 || *b == 6 || *c == 6 || *d == 6 || *e == 6 || *f == 6 || *g == 6 || *h == 6 || *i == 6 {
		found6 = true
		foundCount++
	}

	if *a == 7 || *b == 7 || *c == 7 || *d == 7 || *e == 7 || *f == 7 || *g == 7 || *h == 7 || *i == 7 {
		found7 = true
		foundCount++
	}

	if *a == 8 || *b == 8 || *c == 8 || *d == 8 || *e == 8 || *f == 8 || *g == 8 || *h == 8 || *i == 8 {
		found8 = true
		foundCount++
	}

	if *a == 9 || *b == 9 || *c == 9 || *d == 9 || *e == 9 || *f == 9 || *g == 9 || *h == 9 || *i == 9 {
		found9 = true
		foundCount++
	}

	if foundCount == 8 {
		var numberToFill int

		if !found1 {
			numberToFill = 1
		}
		if !found2 {
			numberToFill = 2
		}
		if !found3 {
			numberToFill = 3
		}
		if !found4 {
			numberToFill = 4
		}
		if !found5 {
			numberToFill = 5
		}
		if !found6 {
			numberToFill = 6
		}
		if !found7 {
			numberToFill = 7
		}
		if !found8 {
			numberToFill = 8
		}
		if !found9 {
			numberToFill = 9
		}

		if *a < 1 {
			*a = numberToFill
		}
		if *b < 1 {
			*b = numberToFill
		}
		if *c < 1 {
			*c = numberToFill
		}
		if *d < 1 {
			*d = numberToFill
		}
		if *e < 1 {
			*e = numberToFill
		}
		if *f < 1 {
			*f = numberToFill
		}
		if *g < 1 {
			*g = numberToFill
		}
		if *h < 1 {
			*h = numberToFill
		}
		if *i < 1 {
			*i = numberToFill
		}
	}

}
