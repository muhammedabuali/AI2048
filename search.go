package AI2048

type Node interface {
	get_parent() Node
	get_operator() int
	get_depth() int
	get_cost() int
	get_children() []Node
}

type Problem interface {
	init_state() Node
	test_success(n Node)
	do_action(x int) Node
	get_cost(n Node, x int)
	get_actions() int //to get the number of childeren b
}

type Strategy func(nodes []Node, children []Node) []Node

func general_search(p Problem, quing_fun Strategy) {

}
