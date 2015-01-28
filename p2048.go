package main

type P2048 struct {
	goal int
	grid Grid // initial from which we create the initial state
}

var global_problem Problem

// Returns the initial state of the board
// As in the Root node of the search tree.
// seed is the seed used by the PRNG.
func (p *P2048) initial_state() Node {
	return &N2048{p.grid, 2, nil, START, 0, 0, 0, false}
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
		new_node := node.apply(LEFT)
		if !gobal_hash[(new_node.(*N2048)).board] {
			nodes = append(nodes, new_node)
			gobal_hash[(new_node.(*N2048)).board] = true
		}
	}
	if node.can_apply(RIGHT) {
		new_node := node.apply(RIGHT)
		if !gobal_hash[(new_node.(*N2048)).board] {
			nodes = append(nodes, new_node)
			gobal_hash[(new_node.(*N2048)).board] = true
		}
	}
	if node.can_apply(DOWN) {
		new_node := node.apply(DOWN)
		if !gobal_hash[(new_node.(*N2048)).board] {
			nodes = append(nodes, new_node)
			gobal_hash[(new_node.(*N2048)).board] = true
		}
	}
	if node.can_apply(UP) {
		new_node := node.apply(UP)
		if !gobal_hash[(new_node.(*N2048)).board] {
			nodes = append(nodes, new_node)
			gobal_hash[(new_node.(*N2048)).board] = true
		}
	}

	return nodes
}
