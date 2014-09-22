package main

type Node interface {
	get_parent() Node
	get_operator() int
	get_depth() int
	get_path_cost() int
}

type Problem interface {
	init_state() Node
	test_success(n Node) bool
	get_action_cost(n Node, x int) int
	expand(n Node) []Node
}

type Strategy func(nodes []Node, children []Node) []Node

func general_search(p Problem, quing_fun Strategy) Node {
	var nodes []Node
	var output Node
	nodes = append(nodes, p.init_state())
	for {
		if len(nodes) == 0 {
			return output
		} else {
			node := nodes[0]
			nodes = nodes[1:]
			if p.test_success(node) {
				return node
			} else {
				var children []Node = p.expand(node)
				nodes = quing_fun(nodes, children)
			}
		}
	}
}
