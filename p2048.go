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

// Returns the node resulting from applying operator direction
// on node n
func make_node(n Node, direction int) *N2048 {
	var node = n.(*N2048)
	b := node.board
	var shift int
	var out *N2048

	if direction == UP {
		for col := 0; col < 4; col++ {
			shift = 0
			//move over empty cells
			for row := 0; row < 4; row++ {
				if b[row][col] == 0 { //empty cell
					shift -= 1
				} else if shift != 0 {
					b[row+shift][col] = b[row][col]
					b[row][col] = 0
				}
			}
			shift = 0
			//sum equal neighbours
			for row := 0; row < 3; row++ {
				if b[row][col] == 0 { //empty cell
					break // no more cells
				}
				if b[row][col] == b[row+1][col] {
					b[row][col] *= 2
					b[row+1][col] = 0
					if shift != 0 {
						b[row+shift][col] = b[row][col]
						b[row][col] = 0
					}
					shift -= 1
					row += 1
				} else if shift != 0 {
					b[row+shift][col] = b[row][col]
					b[row][col] = 0
				}
			}
			if shift != 0 {
				b[3+shift][col] = b[3][col]
				b[3][col] = 0
			}
		}

		// calculate max
		maximum := max_of_grid(&b)
		add_tile(&b)
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == DOWN {
		for col := 0; col < 4; col++ {
			shift = 0
			//move over empty cells
			for row := 3; row > -1; row-- {
				if b[row][col] == 0 { //empty cell
					shift += 1
				} else if shift != 0 {
					b[row+shift][col] = b[row][col]
					b[row][col] = 0
				}
			}
			shift = 0
			//sum equal neighbours
			for row := 3; row > 0; row-- {
				if b[row][col] == 0 { //empty cell
					break // no more cells
				}
				if b[row][col] == b[row-1][col] {
					b[row][col] *= 2
					b[row-1][col] = 0
					if shift != 0 {
						b[row+shift][col] = b[row][col]
						b[row][col] = 0
					}
					shift += 1
					row -= 1
				} else if shift != 0 {
					b[row+shift][col] = b[row][col]
					b[row][col] = 0
				}
			}
			if shift != 0 {
				b[0+shift][col] = b[0][col]
				b[0][col] = 0
			}
		}
		maximum := max_of_grid(&b)
		add_tile(&b)
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == LEFT { //left
		for row := 0; row < 4; row++ {
			shift = 0
			//move over empty cells
			for col := 0; col < 4; col++ {
				if b[row][col] == 0 { //empty cell
					shift -= 1
				} else if shift != 0 {
					b[row][col+shift] = b[row][col]
					b[row][col] = 0
				}
			}
			shift = 0
			//sum equal neighbours
			for col := 0; col < 3; col++ {
				if b[row][col] == 0 { //empty cell
					break // no more cells
				}
				if b[row][col] == b[row][col+1] {
					b[row][col] *= 2
					b[row][col+1] = 0
					if shift != 0 {
						b[row][col+shift] = b[row][col]
						b[row][col] = 0
					}
					shift -= 1
					col += 1
				} else if shift != 0 {
					b[row][col+shift] = b[row][col]
					b[row][col] = 0
				}
			}
			if shift != 0 {
				b[row][3+shift] = b[row][3]
				b[row][3] = 0
			}
		}
		maximum := max_of_grid(&b)
		add_tile(&b)
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == RIGHT {
		for row := 0; row < 4; row++ {
			shift = 0
			//move over empty cells
			for col := 3; col > -1; col-- {
				if b[row][col] == 0 { //empty cell
					shift += 1
				} else if shift != 0 {
					b[row][col+shift] = b[row][col]
					b[row][col] = 0
				}
			}
			shift = 0
			//sum equal neighbours
			for col := 3; col > 0; col-- {
				if b[row][col] == 0 { //empty cell
					break // no more cells
				}
				if b[row][col] == b[row][col-1] {
					b[row][col] *= 2
					b[row][col-1] = 0
					if shift != 0 {
						b[row][col+shift] = b[row][col]
						b[row][col] = 0
					}
					shift += 1
					col -= 1
				} else if shift != 0 {
					b[row][col+shift] = b[row][col]
					b[row][col] = 0
				}
			}
			if shift != 0 {
				b[row][0+shift] = b[row][0]
				b[row][0] = 0
			}
		}
		maximum := max_of_grid(&b)
		add_tile(&b)
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	}
	return out
}

// adds a 2 at the first free corner if there is one
func add_tile(g *Grid) {
	r, c, ok := first_empty_corner(g)
	if ok {
		g[r][c] = 2
	}
}
