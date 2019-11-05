package algo

// InsertSort - O(N^2)
func InsertSort(E []Element) {

	for i := 1; i < len(E); i = i + 1 {

		j := i - 1
		temp := E[i]

		for j >= 0 && temp.LessThan(E[j]) {
			E[j+1] = E[j]
			j = j - 1
		}
		E[j+1] = temp

	}

}

func partition(E []Element, pivot Element) int {

	l, r := 0, len(E)-1

	for i := range E {
		if !E[i].LessThan(pivot) && !pivot.LessThan(E[i]) {
			E[i], E[r] = E[r], E[i]
			break
		}
	}

	for j := range E {
		if E[j].LessThan(pivot) {
			E[l], E[j] = E[j], E[l]
			l++
		}
	}

	E[l], E[r] = E[r], E[l]
	return l

}
