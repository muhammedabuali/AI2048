package main

import (
	"math/rand"
)

type P2048 struct {
	goal int
}

type N2048 struct {
	board     [4][4]int
	max       int
	parent    *N2048
	operator  int
	path_cost int
	depth     int
}

func (*P2048) init_state() Node {
	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	var i, j int
	i = rand.Intn(16)
	j = rand.Intn(16)
	for i == j {
		j = rand.Intn(16)
	}
	board[i/4][i%4] = 2
	board[j/4][j%4] = 2
	var root *N2048
	var node N2048 = N2048{board, 2, root, 0, 0, 0}
	return node
}

func (p *P2048) test_success(n Node) bool {
	var node N2048 = n.(N2048)
	if p.goal == node.max {
		return true
	} else {
		return false
	}
}

func (*P2048) get_action_cost(node Node, x int) int {
	return 1
}

func (*P2048) expand(node Node) []Node {
	var nodes []Node
	for i := 1; i <= 4; i++ {
		var node Node = make_node(node.(N2048), i)
		nodes = append(nodes, node)
	}
	return nodes
}

func make_node(n Node, direction int) N2048 {
	var node = n.(N2048)
	b := node.board
	var shift int
	var out N2048
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
			for row := 0; row < 4; row++ {
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
		}
		//TODO: add 2 @ random position
		// calculate max
		var empty_cells []int
		var maximum int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if len(empty_cells) != 0 {
			pos := rand.Intn(len(empty_cells))
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = N2048{b, maximum, &node, 1, node.path_cost + 1,
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
			for row := 3; row > -1; row-- {
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
		}
		var empty_cells []int
		var maximum int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if len(empty_cells) != 0 {
			pos := rand.Intn(len(empty_cells))
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = N2048{b, maximum, &node, 1, node.path_cost + 1,
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
			for col := 0; col < 4; col++ {
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
		}
		var empty_cells []int
		var maximum int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if len(empty_cells) != 0 {
			pos := rand.Intn(len(empty_cells))
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = N2048{b, maximum, &node, 1, node.path_cost + 1,
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
			for col := 3; col > -1; col-- {
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
		}
		var empty_cells []int
		var maximum int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if b[i][j] == 0 {
					temp := i*4 + j
					empty_cells = append(empty_cells, temp)
				} else if b[i][j] > maximum {
					maximum = b[i][j]
				}
			}
		}
		if len(empty_cells) != 0 {
			pos := rand.Intn(len(empty_cells))
			pos = empty_cells[pos]
			b[pos/4][pos%4] = 2
		}
		out = N2048{b, maximum, &node, 1, node.path_cost + 1,
			node.depth + 1}
	}
	return out
}

func (node N2048) get_parent() Node {
	return node.parent
}

func (node N2048) get_operator() int {
	return node.operator
}

func (node N2048) get_path_cost() int {
	return node.path_cost
}

func (node N2048) get_depth() int {
	return node.depth
}
