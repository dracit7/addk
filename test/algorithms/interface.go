package algorithms

import "github.com/dracit7/algorithms/algo"

type element int

func (e element) LessThan(f algo.Element) bool {
	return e < f.(element)
}
