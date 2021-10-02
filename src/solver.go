package src

import "fmt"

func (b *Board) Solve() bool {
	var didSomething bool = false
	fmt.Printf("[All-but-one] Checking Rows...")
	retVal := b.checkRowsForAllButOne()
	if retVal == "" {
		fmt.Printf("done. (no action) \n")
	} else {
		fmt.Printf("done.  %s \n", retVal)
		didSomething = true
	}

	fmt.Printf("[All-but-one] Checking Columns...")
	retVal = b.checkColumnsForAllButOne()
	if retVal == "" {
		fmt.Printf("done. (no action) \n")
	} else {
		fmt.Printf("done.  %s \n", retVal)
		didSomething = true
	}

	fmt.Printf("[All-but-one] Checking Sectors...")
	retVal = b.checkSectorsForAllButOne()
	if retVal == "" {
		fmt.Printf("done. (no action) \n")
	} else {
		fmt.Printf("done.  %s \n", retVal)
		didSomething = true
	}

	fmt.Printf("[Crossfill-By-Sector] Checking Crossfills...")
	retVal = b.checkCrossFillBySector()
	if retVal == "" {
		fmt.Printf("done. (no action) \n")
	} else {
		fmt.Printf("done.  %s \n", retVal)
		didSomething = true
	}


	return didSomething
}
