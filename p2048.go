package main

type P2048 struct {
	goal int
}

type N2048 struct {
	board [4][4]int
	max   int
}

func (*P2048) init_state() N2048 {
	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	var node N2048 = N2048{board, 0}
	return node
}

func (p *P2048) test_success(node N2048) bool {
	if p.goal == node.max {
		return true
	} else {
		return false
	}
}

func (*P2048) get_action_cost(node N2048, x int) int {
	return 1
}

func (*P2048) expand(node N2048) []N2048 {
	var nodes []N2048
	for i := 1; i <= 4; i++ {
		var node N2048 = make_node(node, i)
		nodes = append(nodes, node)
	}
	return nodes
}

func make_node(node N2048, direction int) N2048 {
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
		out = N2048{b, 0}
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
		out = N2048{b, 0}
	}

	return out
}
