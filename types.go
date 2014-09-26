package AI2048

// Operator/Direction constants
const (
	LEFT  = iota
	RIGHT = iota
	UP    = iota
	DOWN  = iota
	START = iota // operator value of root Node
)

type Node interface {
	get_parent() Node
	// Returns the operator used to reach this node
	get_operator() int
	// Returns path depth with root having a depth of 0
	get_depth() int
	// Returns the path cost, which represents the score
	get_path_cost() int
	// Returns the path from root to this node (series of actions)
	get_path() Path
	// Returns true if it's possible to apply operator on this node
	can_apply(operator int) bool
	// Applies operator and returns the resulting node
	apply(operator int) Node
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
