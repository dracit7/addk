package algo

// Item defines the item you want to
// put in the knapsack.
type Item interface {
	Weight() int
	Value() int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Knapsack01 solves the 0-1 knapsack problem.
func Knapsack01(I []Item, size int) int {

	// Make a matrix to store interim results
	// of the DP algorithm.
	//
	// dpMem[i][sizeleft] stores the optimal
	// total value with first `i` items and
	// `sizeleft` space.
	dpMem := make([][]int, len(I)+1)
	for i := 0; i <= len(I); i++ {
		dpMem[i] = make([]int, size+1)
	}

	// For the first `i` items:
	for i, item := range I {
		// If the space left is `sizeLeft`:
		for sizeLeft := 1; sizeLeft <= size; sizeLeft++ {
			// Calculate the optimum total value in this occasion.
			if sizeLeft >= item.Weight() {
				// There's enough space for item `i`, we record
				// the higher total value between picking item `i`
				// and the max value when we abandon item `i`.
				dpMem[i+1][sizeLeft] = max(dpMem[i][sizeLeft-item.Weight()]+item.Value(), dpMem[i][sizeLeft])
			} else {
				dpMem[i+1][sizeLeft] = dpMem[i][sizeLeft]
			}
		}
	}

	return dpMem[len(I)][size]

}
