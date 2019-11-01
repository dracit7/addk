package graph

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

const (
	// UNDIRECTED represents undirected graph
	UNDIRECTED = 0
	// DIRECTED represents directed graph
	DIRECTED = 1
)

/**********************************
 *        Adjacent Table          *
 **********************************/

// Verticle *
type Verticle struct {

	// Global member
	ID     interface{} `yaml:"id"`
	Serial int

	// Weighted graph, represented by adjacent table
	Edges []Edge `yaml:"edges"`
}

// Edge *
type Edge struct {
	Endpoint   int         `yaml:"-"`        // endpoint's serial
	EndpointID interface{} `yaml:"endpoint"` // endpoint's ID
	Weight     int         `yaml:"weight"`
}

// Graph *
type Graph struct {
	Type int        `yaml:"type"`
	V    []Verticle `yaml:"v"`
	E    []SEdge
}

/**********************************
 *            Nodes               *
 **********************************/

// Node *
type Node struct {
	ID       string
	Neighbor []*Node
}

/**********************************
 *            Edges               *
 **********************************/

// SEdge contains two endpoint's info
// instead of one.
type SEdge struct {
	Weight    int
	EndpointA int
	EndpointB int
}

/**********************************
 *           Functions            *
 **********************************/

/*					Verticles							*/

// NewVerticle returns a new Verticle initialized by `conf`.
func NewVerticle(conf interface{}) (*Verticle, error) {

	var v Verticle
	var err error

	// We support various initializing methods.
	switch conf.(type) {

	case string:
		err = yaml.Unmarshal([]byte(conf.(string)), &v)
		if err != nil {
			return nil, err
		}

	case []byte:
		err = yaml.Unmarshal(conf.([]byte), &v)
		if err != nil {
			return nil, err
		}

	}

	// We do not deal with v.Endpoint here,
	// because there's no chance to know the serial of
	// the verticle with id == v.EndpointID in graph.
	//
	// The translation from EndpointID to Endpoint
	// takes place in `AddVerticle()`.
	return &v, nil

}

// This should not be used by user because it's unsafe.
//
// We must update some other essential states of this
// graph after calling this.
func (v *Verticle) addEdge(e *Edge) error {

	v.Edges = append(v.Edges, *e)
	return nil

}

/*					Graphs							*/

// NewGraph returns a new graph initialized by `conf`.
func NewGraph(conf interface{}) (*Graph, error) {

	var g Graph
	var err error

	// We support various initializing methods.
	switch conf.(type) {

	case string:
		err = yaml.Unmarshal([]byte(conf.(string)), &g)
		if err != nil {
			return nil, err
		}

	case []byte:
		err = yaml.Unmarshal(conf.([]byte), &g)
		if err != nil {
			return nil, err
		}

	case []Verticle:
		g.V = conf.([]Verticle)

	}

	// Do some setups after Unmarshal().

	// Set the serial of each verticle.
	for i := range g.V {
		g.V[i].Serial = i
	}

	// We use Serial as endpoint instead of ID for efficiency,
	// but use ID to design graph for readability.
	//
	// So we need to do a translation here.
	for i := range g.V {
		for j, e := range g.V[i].Edges {
			g.V[i].Edges[j].Endpoint = g.GetVerticleByID(e.EndpointID).Serial
		}
	}

	// Translate V to E
	for _, v := range g.V {
		for _, e := range v.Edges {
			// If `g` is an undirected graph, to avoid replicated
			// edges, we only record half of edges.
			if v.Serial < e.Endpoint || g.Type == DIRECTED {
				g.E = append(g.E, SEdge{
					Weight:    e.Weight,
					EndpointA: v.Serial,
					EndpointB: e.Endpoint,
				})
			}
		}
	}

	// Now each edge's endpoint is equal to relevant verticle's index.

	return &g, nil

}

// GetVerticleByID finds a verticle in `g` with specified id.
func (g *Graph) GetVerticleByID(id interface{}) *Verticle {

	for _, v := range g.V {
		if v.ID == id {
			return &v
		}
	}

	return nil

}

// AddVerticle adds a new verticle to `g`.
func (g *Graph) AddVerticle(v *Verticle) error {

	// Append v to g.
	v.Serial = len(g.V)
	g.V = append(g.V, *v)

	// Translate IDs if necessary
	for i, e := range v.Edges {
		v.Edges[i].Endpoint = g.GetVerticleByID(e.EndpointID).Serial
		g.E = append(g.E, SEdge{
			Weight:    e.Weight,
			EndpointA: v.Serial,
			EndpointB: v.Edges[i].Endpoint,
		})
	}

	// For undirected graph, update connection status of other verticles.
	if g.Type == UNDIRECTED {
		for _, e := range v.Edges {
			g.V[e.Endpoint].addEdge(&Edge{
				EndpointID: v.ID,
				Endpoint:   v.Serial,
				Weight:     e.Weight,
			})
		}
	}

	return nil

}

// Connect creates an edge from srcID to destID with weight.
func (g *Graph) Connect(srcID interface{}, destID interface{}, weight int) error {

	// Get verticles
	src := g.GetVerticleByID(srcID)
	dst := g.GetVerticleByID(destID)
	if src == nil || dst == nil {
		return fmt.Errorf("No such verticle in this graph")
	}

	// Add edge
	err := g.V[src.Serial].addEdge(&Edge{
		Weight:     weight,
		EndpointID: destID,
		Endpoint:   dst.Serial,
	})
	if err != nil {
		return err
	}

	g.E = append(g.E, SEdge{
		Weight:    weight,
		EndpointA: src.Serial,
		EndpointB: dst.Serial,
	})

	// For undirected graph, update peer's connection state.
	if g.Type == UNDIRECTED {
		err = g.V[dst.Serial].addEdge(&Edge{
			Weight:     weight,
			EndpointID: srcID,
			Endpoint:   src.Serial,
		})
		if err != nil {
			return err
		}
	}

	return nil

}

// ToNodes transfer a graph into nodes representation.
func (g *Graph) ToNodes() []*Node {

	nodes := make([]*Node, len(g.V))

	// Initialize all nodes
	for i := range g.V {
		nodes[i] = &Node{
			ID:       fmt.Sprintf("%+v", g.V[i].ID),
			Neighbor: []*Node{},
		}
	}

	// Intepret relations
	for i, v := range g.V {
		for _, e := range v.Edges {
			nodes[i].Neighbor = append(nodes[i].Neighbor, nodes[e.Endpoint])
		}
	}

	return nodes

}

// ToEdges transfer a graph into edges representation.
func (g *Graph) ToEdges() []SEdge {
	return g.E
}
