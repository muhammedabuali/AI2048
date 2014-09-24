package AI2048

import "math/rand"

type P2048 struct {
	goal int
	grid *Grid
}

const (
	LEFT  = iota
	RIGHT = iota
	UP    = iota
	DOWN  = iota
	START = iota
)

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

// All actions have a cost of one.
func (*P2048) get_action_cost(node Node, x int) int {
	return 1
}

// Returns all children of a Node
func (*P2048) expand(node Node) []Node {
	nodes := make([]Node, 0, 4)
	// Branching factor is four (four operators)
	for i := 1; i <= 4; i++ {
		node := make_node(node.(*N2048), i)
		nodes = append(nodes, node)
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
	if direction == 1 {
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
		//TODO: add 2 @ random position
		// calculate max
		var empty_cells []int
		var maximum, count int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
					count += 1
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if count != 0 {
			pos := rand.Intn(count)
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == 2 {
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
		var empty_cells []int
		var maximum, count int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
					count += 1
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if count != 0 {
			pos := rand.Intn(count)
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == 3 { //left
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
		var empty_cells []int
		var maximum, count int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
					count += 1
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if count != 0 {
			pos := rand.Intn(count)
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	} else if direction == 4 { //right
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
		var empty_cells []int
		var maximum, count int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
					count += 1
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if count != 0 {
			pos := rand.Intn(count)
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = &N2048{b, maximum, node, 1, node.path_cost + 1,
			node.depth + 1}
	}
	return out
}

func add_tile(g *Grid, value int) {

}

// Returns 3-tuple (r, c, ok) where r is row number c is column number
// and ok is true if there is an empty corner
// Corner order : Tope-left and then clock-wise
func first_empty_corner(g *Grid) (r, c int, ok bool) {
	if g[0][0] == 0 {
		return 0, 0, true
	} else if g[0][3] == 0 {
		return 0, 3, true
	} else if g[3][3] == 0 {
		return 3, 3, true
	} else if g[3][0] == 0 {
		return 3, 0, true
	} else {
		return -1, -1, false
	}
}
