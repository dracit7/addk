package algo

import "fmt"

// MatrixChainMultiOrder returns the optimal order for chain
// multiplying matrices.
func MatrixChainMultiOrder(dims []int) string {

	n := len(dims) - 1

	// dpMem[i][j] will be the minimum cost of the multiplication
	// A[i]*A[i+1]*...*A[j]. Be cautious that dpMem[i][i] = 0 (no cost).
	dpMem := newMatrix(n)
	// bestK[i][j] will be the index of the subsequence split that
	// achieved minimal cost.
	bestK := newMatrix(n)

	// Assume that `k` is the optimal separation index which
	// minimizes the total cost. i.e. A[0...k]*A[k...n] minimizes
	// the total cost.
	//
	// Let `l` be the distance between i&j, then:
	for l := 1; l < n; l++ {

		// Move a "window" with length `l` from 0~l to (n-l)~n.
		for i := 0; i < n-l; i++ {
			j := i + l

			// Mark dpMem[i][j] as -1 at first
			dpMem[i][j] = -1

			// Try each possible k between i and j
			// and calculate the relevant dpMem.
			for k := i; k < j; k++ {
				cost := dpMem[i][k] + dpMem[k+1][j] + dims[i]*dims[k+1]*dims[j+1]

				// Update dpMem and bestK if we find a better k.
				if dpMem[i][j] == -1 || cost < dpMem[i][j] {
					dpMem[i][j] = cost
					bestK[i][j] = k
				}
			}
		}
	}

	// Use a recursive function to format the
	// result into a string.
	var formatter func(i int, j int) string

	formatter = func(i int, j int) string {
		if i == j {
			return fmt.Sprintf("%d", i)
		}
		k := bestK[i][j]
		return fmt.Sprintf("(%s)*(%s)", formatter(i, k), formatter(k+1, j))
	}

	return formatter(0, n-1)

}

// Returns a n*n matrix to store interim results
// of the DP algorithm.
func newMatrix(n int) (m [][]int) {
	m = make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	return m
}
