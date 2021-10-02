package src2

func buildBoard1() Board {
	c := make([]Cell, 0)


	type Cell struct {
	X       int
	Y       int
	Filled  bool
	Address string
	Value   int
}


	return Board{
		A2: 6,
		A4: 9,
		A5: 8,
		A6: 3,
		A8: 7,
		B2: 9,
		B5: 1,
		B7: 6,
		B8: 2,
		C3: 4,
		C6: 6,
		C9: 3,
		D1: 2,
		F2: 4,
		F7: 7,
		G4: 2,
		G6: 1,
		G7: 3,
		H1: 1,
		H2: 7,
		H6: 8,
		H8: 9,
		H9: 2,
		I5: 3,
		I7: 4,
		I8: 6,
	}
}


}
