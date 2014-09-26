package main

import (
	"math"
)

func general_search(p Problem, quing_fun Strategy) (Node, bool, uint64) {
	nodes := make([]Node, 0, 10)             // Make a queue
	nodes = append(nodes, p.initial_state()) // queue initial state
	expanded_nodes := uint64(0)
	for {
		if len(nodes) == 0 {
			return nil, false, expanded_nodes
		} else {
			// Remove first node
			node := nodes[0]
			nodes = nodes[1:]
			if p.goal_test(node) {
				return node, true, expanded_nodes
			} else {
				expanded_nodes++
				nodes = quing_fun(nodes, p.expand(node))
			}
		}
	}
}

func iterative_deepening_search(p Problem) (Node, bool, uint64) {

	total_expanded_nodes := uint64(0)
	for limit := uint64(0); limit < math.MaxUint64; limit++ {
		quing_fun := depth_limited_search(limit)
		target, success, expanded_nodes := general_search(p, quing_fun)
		total_expanded_nodes += expanded_nodes
		if success {
			return target, success, total_expanded_nodes
		}

	}

	return nil, false, total_expanded_nodes
}
