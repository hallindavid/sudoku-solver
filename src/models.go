package solver

type Board struct {
	A1 int
	A2 int
	A3 int
	A4 int
	A5 int
	A6 int
	A7 int
	A8 int
	A9 int

	B1 int
	B2 int
	B3 int
	B4 int
	B5 int
	B6 int
	B7 int
	B8 int
	B9 int

	C1 int
	C2 int
	C3 int
	C4 int
	C5 int
	C6 int
	C7 int
	C8 int
	C9 int

	D1 int
	D2 int
	D3 int
	D4 int
	D5 int
	D6 int
	D7 int
	D8 int
	D9 int

	E1 int
	E2 int
	E3 int
	E4 int
	E5 int
	E6 int
	E7 int
	E8 int
	E9 int

	F1 int
	F2 int
	F3 int
	F4 int
	F5 int
	F6 int
	F7 int
	F8 int
	F9 int

	G1 int
	G2 int
	G3 int
	G4 int
	G5 int
	G6 int
	G7 int
	G8 int
	G9 int

	H1 int
	H2 int
	H3 int
	H4 int
	H5 int
	H6 int
	H7 int
	H8 int
	H9 int

	I1 int
	I2 int
	I3 int
	I4 int
	I5 int
	I6 int
	I7 int
	I8 int
	I9 int
}

func (b Board) SolvedCount() int {
	return 81 - b.UnsolvedCount()
}

func (b Board) UnsolvedCount() int {
	counter := 81

	if b.A1 > 0 {
		counter--
	}
	if b.A2 > 0 {
		counter--
	}
	if b.A3 > 0 {
		counter--
	}
	if b.A4 > 0 {
		counter--
	}
	if b.A5 > 0 {
		counter--
	}
	if b.A6 > 0 {
		counter--
	}
	if b.A7 > 0 {
		counter--
	}
	if b.A8 > 0 {
		counter--
	}
	if b.A9 > 0 {
		counter--
	}

	if b.B1 > 0 {
		counter--
	}
	if b.B2 > 0 {
		counter--
	}
	if b.B3 > 0 {
		counter--
	}
	if b.B4 > 0 {
		counter--
	}
	if b.B5 > 0 {
		counter--
	}
	if b.B6 > 0 {
		counter--
	}
	if b.B7 > 0 {
		counter--
	}
	if b.B8 > 0 {
		counter--
	}
	if b.B9 > 0 {
		counter--
	}

	if b.C1 > 0 {
		counter--
	}
	if b.C2 > 0 {
		counter--
	}
	if b.C3 > 0 {
		counter--
	}
	if b.C4 > 0 {
		counter--
	}
	if b.C5 > 0 {
		counter--
	}
	if b.C6 > 0 {
		counter--
	}
	if b.C7 > 0 {
		counter--
	}
	if b.C8 > 0 {
		counter--
	}
	if b.C9 > 0 {
		counter--
	}

	if b.D1 > 0 {
		counter--
	}
	if b.D2 > 0 {
		counter--
	}
	if b.D3 > 0 {
		counter--
	}
	if b.D4 > 0 {
		counter--
	}
	if b.D5 > 0 {
		counter--
	}
	if b.D6 > 0 {
		counter--
	}
	if b.D7 > 0 {
		counter--
	}
	if b.D8 > 0 {
		counter--
	}
	if b.D9 > 0 {
		counter--
	}

	if b.E1 > 0 {
		counter--
	}
	if b.E2 > 0 {
		counter--
	}
	if b.E3 > 0 {
		counter--
	}
	if b.E4 > 0 {
		counter--
	}
	if b.E5 > 0 {
		counter--
	}
	if b.E6 > 0 {
		counter--
	}
	if b.E7 > 0 {
		counter--
	}
	if b.E8 > 0 {
		counter--
	}
	if b.E9 > 0 {
		counter--
	}

	if b.F1 > 0 {
		counter--
	}
	if b.F2 > 0 {
		counter--
	}
	if b.F3 > 0 {
		counter--
	}
	if b.F4 > 0 {
		counter--
	}
	if b.F5 > 0 {
		counter--
	}
	if b.F6 > 0 {
		counter--
	}
	if b.F7 > 0 {
		counter--
	}
	if b.F8 > 0 {
		counter--
	}
	if b.F9 > 0 {
		counter--
	}

	if b.G1 > 0 {
		counter--
	}
	if b.G2 > 0 {
		counter--
	}
	if b.G3 > 0 {
		counter--
	}
	if b.G4 > 0 {
		counter--
	}
	if b.G5 > 0 {
		counter--
	}
	if b.G6 > 0 {
		counter--
	}
	if b.G7 > 0 {
		counter--
	}
	if b.G8 > 0 {
		counter--
	}
	if b.G9 > 0 {
		counter--
	}

	if b.H1 > 0 {
		counter--
	}
	if b.H2 > 0 {
		counter--
	}
	if b.H3 > 0 {
		counter--
	}
	if b.H4 > 0 {
		counter--
	}
	if b.H5 > 0 {
		counter--
	}
	if b.H6 > 0 {
		counter--
	}
	if b.H7 > 0 {
		counter--
	}
	if b.H8 > 0 {
		counter--
	}
	if b.H9 > 0 {
		counter--
	}

	if b.I1 > 0 {
		counter--
	}
	if b.I2 > 0 {
		counter--
	}
	if b.I3 > 0 {
		counter--
	}
	if b.I4 > 0 {
		counter--
	}
	if b.I5 > 0 {
		counter--
	}
	if b.I6 > 0 {
		counter--
	}
	if b.I7 > 0 {
		counter--
	}
	if b.I8 > 0 {
		counter--
	}
	if b.I9 > 0 {
		counter--
	}
	return counter

}
