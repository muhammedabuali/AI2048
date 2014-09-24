package AI2048

type Node interface {
	get_parent() Node
	get_operator() int
	get_depth() int
	get_path_cost() int
	get_path() Path
}

type Path []string

type Problem interface {
	initial_state() Node
	goal_test(n Node) bool
	get_action_cost(n Node, x int) int
	expand(n Node) []Node
}

type Strategy func(nodes []Node, children []Node) []Node

type heuristic func(Node) int
