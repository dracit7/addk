package algorithms

import "github.com/dracit7/algorithms/algo"

type element int

func (e element) LessThan(f algo.Element) bool {
	return e < f.(element)
}

type item struct {
	value  int
	weight int
}

func (i item) Value() int {
	return i.value
}

func (i item) Weight() int {
	return i.weight
}
