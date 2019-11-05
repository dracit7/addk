package algo

// InsertSort - O(N^2)
func InsertSort(E *[]Element) {

	for i := 1; i < len(*E); i = i + 1 {

		j := i - 1
		temp := (*E)[i]

		for j >= 0 && temp.LessThan((*E)[j]) {
			(*E)[j+1] = (*E)[j]
			j = j - 1
		}
		(*E)[j+1] = temp

	}

}

func partition(E []Element, pivot Element) int {

	lptr, rptr := 1, len(E)-1

	for {

		// Find a swappable element on each side of E
		for lptr < rptr && E[lptr].LessThan(pivot) {
			lptr = lptr + 1
		}
		for rptr >= 0 && pivot.LessThan(E[rptr]) {
			rptr = rptr - 1
		}

		// Swap them. If lptr has met rptr, return rptr.
		if lptr < rptr {
			E[lptr], E[rptr] = E[rptr], E[lptr]
		} else {
			return rptr
		}

	}

}
