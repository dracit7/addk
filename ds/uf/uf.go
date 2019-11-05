package uf

// Node is the unit of a Union-Find set.
//
// Each Node keeps a pointer which points to
// its parent in the uf set. The root's `Parent`
// is nil.
//
// Since we only need to trace the root of an uf
// set, there's no need to add any other pointer.
type Node struct {
	ID     int
	Parent *Node
}

// NewUFSet returns a uf set with a bare root.
func NewUFSet(id int) *Node {
	n := &Node{
		ID:     id,
		Parent: nil,
	}
	n.Parent = n
	return n
}

// Findroot finds the root of the uf set which contains `n`.
func Findroot(n *Node) *Node {
	for {
		if n.Parent == n {
			return n
		}
		n = n.Parent
	}
}

// Union combines to union-find sets.
func Union(ufa *Node, ufb *Node) {
	ufb.Parent = ufa
}
