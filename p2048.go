package main

type P2048 struct {
	goal int
	grid *Grid
}

// Returns the initial state of the board
// As in the Root node of the search tree.
// seed is the seed used by the PRNG.
func (p *P2048) initial_state() Node {
	return &N2048{*p.grid, 2, nil, START, 0, 0}
}

// Returns true if we have reached M.
func (p *P2048) goal_test(n Node) bool {
	return p.goal == n.(*N2048).max
}

// Returns all children of a Node
func (*P2048) expand(node Node) []Node {
	nodes := make([]Node, 0, 4)
	// Apply all possible operators
	if node.can_apply(LEFT) {
		nodes = append(nodes, node.apply(LEFT))
	}
	if node.can_apply(RIGHT) {
		nodes = append(nodes, node.apply(RIGHT))
	}
	if node.can_apply(DOWN) {
		nodes = append(nodes, node.apply(DOWN))
	}
	if node.can_apply(UP) {
		nodes = append(nodes, node.apply(UP))
	}

	return nodes
}
