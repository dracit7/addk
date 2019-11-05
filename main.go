package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dracit7/algorithms/algo"
	"github.com/dracit7/algorithms/ds/graph"
	"github.com/dracit7/algorithms/visual"
)

type node struct {
	ID   int
	Next []*node
}

type element int

func (e element) LessThan(f algo.Element) bool {
	return e < f.(element)
}

func main() {

	s, err := ioutil.ReadFile("input/ds/graph/mst2.yaml")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	g, err := graph.NewGraph(s)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	g.MST()

	spec := &visual.Spec{
		MST: true,
	}

	visual.Dump(g, "mst", spec)

}
