package algo

// Element is any type that supports comparing
type Element interface {
	LessThan(j Element) bool
}

// Min returns minimal elements in `E`.
func Min(E []Element) Element {

	if len(E) == 0 {
		return nil
	}

	min := E[0]
	for _, e := range E {
		if (e.LessThan(min)) {
			min = e
		}
	}

	return min

}