package ds

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/dracit7/algorithms/ds/graph"
	"github.com/dracit7/algorithms/visual"
)

func TestMST(t *testing.T) {

	s, err := ioutil.ReadFile("input/ds/graph/testcase.yaml")
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
