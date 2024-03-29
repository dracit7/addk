package graph

import (
	"container/heap"

	"github.com/dracit7/algorithms/ds/uf"
)

// Implement a heap interface for edges.

type eHeap []SEdge

func (h eHeap) Len() int {
	return len(h)
}

func (h eHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h eHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h *eHeap) Push(e interface{}) {
	*h = append(*h, e.(SEdge))
}

func (h *eHeap) Pop() interface{} {
	n := len(*h)
	e := (*h)[n-1]
	*h = (*h)[:n-1]
	return e
}

// MST returns the minimize spanning tree of `g`.
//
// We use the optimized kruskal's algo which
// reaches the pseudo-linear complexity.
func (g *Graph) MST() *Graph {

	mst, _ := NewGraph("")

	// Initialize all union-find sets
	// This step costs O(V).
	ufs := make([]*uf.Node, len(g.V))
	for i := range g.V {
		ufs[i] = uf.NewUFSet(i)
	}

	// Use a min-heap to optimize the algo.
	//
	// Building a min-heap only costs O(E),
	// and every pop operation costs O(logK) where
	// K is equal to the size of heap.
	EHeap := eHeap(g.E)
	heap.Init(&EHeap)

	for i := 0; i < len(g.V)-1; i++ {

		// Get the shortest edge
		//
		// This step costs O(logK). If we use a
		// soft heap here, the amortized complexity
		// would be O(1) instead.
		e := heap.Pop(&EHeap)

		// Find the union-find sets of each endpoints
		//
		// This step costs O(1) with path compression.
		ufa := uf.Findroot(ufs[e.(SEdge).EndpointA])
		ufb := uf.Findroot(ufs[e.(SEdge).EndpointB])

		// Path compression.
		//
		// This step costs O(α(N))
		ufs[e.(SEdge).EndpointA].Parent = ufa
		ufs[e.(SEdge).EndpointB].Parent = ufb

		// If two verticles are in different ufs,
		// combine their ufs and record the edge to mst.
		if ufa != ufb {

			// Combine two ufs.
			//
			// This step costs O(1).
			uf.Union(ufa, ufb)

			// Check if va and vb are in MST.
			//
			// This step costs O(1) if implemented properly.
			va := g.V[e.(SEdge).EndpointA]
			vb := g.V[e.(SEdge).EndpointB]
			vaExist := mst.GetVerticleByID(va.ID)
			vbExist := mst.GetVerticleByID(vb.ID)

			// Build verticles to be added to mst.
			v1 := &Verticle{
				ID:    va.ID,
				Edges: []Edge{},
			}
			v2 := &Verticle{
				ID:    vb.ID,
				Edges: []Edge{},
			}

			// Ensure that va and vb are added to mst.
			//
			// This step costs O(1).
			if vaExist == nil && vbExist == nil {
				mst.AddVerticle(v1)
				mst.AddVerticle(v2)
			} else if vaExist != nil && vbExist == nil {
				mst.AddVerticle(v2)
			} else if vaExist == nil && vbExist != nil {
				mst.AddVerticle(v1)
			} else {
				// just pass
			}
			mst.Connect(v1.ID, v2.ID, e.(SEdge).Weight)

			// Set va and vb's InMST flags
			//
			// This step costs O(1) in most cases, and can be omitted if
			// we don't want to visualize the data structure.
			for i, edge := range g.V[e.(SEdge).EndpointA].Edges {
				if edge.Endpoint == e.(SEdge).EndpointB {
					g.V[e.(SEdge).EndpointA].Edges[i].InMST = true
				}
			}
			for i, edge := range g.V[e.(SEdge).EndpointB].Edges {
				if edge.Endpoint == e.(SEdge).EndpointA {
					g.V[e.(SEdge).EndpointB].Edges[i].InMST = true
				}
			}

		} else {
			// If ufa==ufb, dispose this edge
			i--
		}

	}

	return mst

}
