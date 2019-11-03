package visual

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/dracit7/algorithms/ds/graph"
)

// Spec stores mapping specifications
type Spec struct {
	MST bool
}

var path = "instances/dot"

// Map transforms a data structure into dot language
func Map(w io.Writer, is interface{}, spec *Spec) error {

	var appendix string

	fmt.Fprintln(w, "strict graph {")
	fmt.Fprintln(w, "  node [shape=\"circle\"];")

	switch is.(type) {
	case *graph.Graph:
		// Initialize nodes
		for _, v := range is.(*graph.Graph).V {
			fmt.Fprintf(w, "  %d [label=\"%+v\"]\n", v.Serial, v.ID)
		}
		// Initialize edges
		for _, e := range is.(*graph.Graph).E {

			// Visualize mst
			if (is.(*graph.Graph).V[e.EndpointA].GetEdgeByEndpoint(e.EndpointB).InMST == true ||
				is.(*graph.Graph).V[e.EndpointB].GetEdgeByEndpoint(e.EndpointA).InMST == true) && spec.MST == true {
				appendix = " color=\"red\" style=\"bold\""
			}

			if is.(*graph.Graph).Type == graph.UNDIRECTED {
				fmt.Fprintf(w, "  %d -- %d [label=\"%d\"%s]\n", e.EndpointA, e.EndpointB, e.Weight, appendix)
			} else {
				fmt.Fprintf(w, "  %d -> %d [label=\"%d\"%s]\n", e.EndpointA, e.EndpointB, e.Weight, appendix)
			}

		}
	// case []graph.SEdge:
	// case []*graph.Node:
	default:
		return fmt.Errorf("Unsupported Type")
	}

	fmt.Fprintln(w, "}")

	return nil

}

// Marshal transform a data structure to visible string.
func Marshal(ds interface{}, spec *Spec) (buf *bytes.Buffer) {
	Map(buf, ds, spec)
	return
}

// Dump the Marshaled text to a .dot file.
func Dump(ds interface{}, name string, spec *Spec) error {

	buf := &bytes.Buffer{}
	Map(buf, ds, spec)

	err := ioutil.WriteFile(path+"/"+name+".dot", buf.Bytes(), 0644)
	return err

}

// SetPath sets the root path of instances
func SetPath(newpath string) {
	path = newpath
}
