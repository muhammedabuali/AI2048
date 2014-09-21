package main

type Node interface {
	get_parent() Node
	get_operator() int
	get_depth() int
	get_path_cost() int
	get_children() []Node
}

type Problem interface {
	init_state() Node
	test_success(n Node) bool
	do_action(x int) Node
	get_action_cost(n Node, x int)
	get_actions() int //to get the number of childeren b
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
				var children []Node = node.get_children()
				nodes = quing_fun(nodes, children)
			}
		}
	}
}
