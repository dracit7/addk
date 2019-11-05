package algo

// KthSmallest returns the kth smallest element
func KthSmallest(E []Element, k int) Element {
	shell := append([]Element{}, E...)
	return kSmallest(shell, k, 5)
}

// KthSmallestCustom returns the kth smallest element,
// and allow user to specify the `partsize` in the median
// of median algo.
func KthSmallestCustom(E []Element, k int, partSize int) Element {
	shell := append([]Element{}, E...)
	return kSmallest(shell, k, partSize)
}

func kSmallest(E []Element, k int, partSize int) Element {

	var M []Element
	var v Element

	// Recursive exit
	if len(E) <= partSize {
		InsertSort(E)
		return E[k-1]
	}

	// Find the median of each part
	for i := 0; i < len(E); i += partSize {
		var m []Element
		for j := i; j < len(E) && j < i+partSize; j++ {
			m = append(m, E[j])
		}
		InsertSort(m)
		M = append(M, m[len(m)/2])
	}

	// Find the median of median
	if len(M) == 1 {
		v = M[0]
	} else {
		v = kSmallest(M, len(M)/2, partSize)
	}

	// Use the median of median to part the list
	var pos = partition(E, v)

	// Recursive part
	if k-1 == pos {
		return v
	} else if k-1 < pos {
		return kSmallest(E[:pos], k, partSize)
	} else {
		return kSmallest(E[pos+1:], k-pos-1, partSize)
	}

}
