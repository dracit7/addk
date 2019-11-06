package algorithms

import (
	"fmt"
	"testing"

	"github.com/dracit7/algorithms/algo"
)

func TestKnapSack01(t *testing.T) {

	fmt.Printf("%d\n", algo.Knapsack01(
		[]algo.Item{
			item{60, 10},
			item{100, 20},
			item{120, 30},
		}, 50,
	))

}
