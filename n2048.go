package AI2048

type Grid [4][4]int

type N2048 struct {
	board     Grid
	max       int    // maximum value in board
	parent    *N2048 // nil if root
	operator  int
	path_cost int
	depth     int
}

func (node *N2048) get_parent() Node {
	return node.parent
}

func (node *N2048) get_operator() int {
	return node.operator
}

func (node *N2048) get_path_cost() int {
	return node.path_cost
}

func (node *N2048) get_depth() int {
	return node.depth
}

func (node *N2048) get_path() Path {
	path := make(Path, node.path_cost)
	curr_node := node
	translation_map := make(map[int]string)
	translation_map[LEFT] = "left"
	translation_map[RIGHT] = "right"
	translation_map[UP] = "up"
	translation_map[DOWN] = "down"
	translation_map[START] = "START"
	for i := node.path_cost - 1; i >= 0; i-- {
		path[i] = translation_map[curr_node.operator]
		curr_node = curr_node.parent
	}

	return path

}
