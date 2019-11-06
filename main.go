package main

import (
	"fmt"

	"github.com/dracit7/algorithms/algo"
)

func main() {

	fmt.Printf("%s\n", algo.MatrixChainMultiOrder([]int{
		1, 5, 25, 30, 100, 70, 2, 1, 100, 250, 1, 1000, 2,
	}))

}
