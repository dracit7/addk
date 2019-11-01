package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dracit7/algorithms/ds/graph"
	"github.com/dracit7/algorithms/visual"
)

type node struct {
	ID   int
	Next []*node
}

func main() {

	s, err := ioutil.ReadFile("test/ds/graph/example.yaml")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	g, err := graph.NewGraph(s)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	mst := g.MST()

	visual.Dump(mst, "mst")
	visual.Dump(g, "g")

}
