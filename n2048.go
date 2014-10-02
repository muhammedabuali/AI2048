package main

import (
	"fmt"
)

type Grid [4][4]int

type N2048 struct {
	board     Grid   // The grid
	max       int    // maximum value in board
	parent    *N2048 // nil if root
	operator  int    // operator applied on parent to reach this node
	path_cost int    // cost from initial state to here
	depth     int    // distance from initial state
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

	path := make(Path, node.depth+1)
	curr_node := node
	translation_map := make(map[int]string)
	translation_map[LEFT] = "left"
	translation_map[RIGHT] = "right"
	translation_map[UP] = "up"
	translation_map[DOWN] = "down"
	translation_map[START] = "START"
	for i := node.depth; i >= 0; i-- {
		fmt.Println(curr_node.board)
		fmt.Println(curr_node.get_path_cost())
		path[i] = translation_map[curr_node.operator]
		curr_node = curr_node.parent
	}
	return path
}

func (this *N2048) can_apply(operator int) bool {

	switch operator {
	case LEFT:
		return this.can_apply_helper(0, -1)
	case RIGHT:
		return this.can_apply_helper(0, 1)
	case DOWN:
		return this.can_apply_helper(1, 0)
	case UP:
		return this.can_apply_helper(-1, 0)
	case START:
		return false
	default:
		panic(fmt.Sprintf(
			"unknown argument passed '%v' to can_apply on N2048:\n%v",
			operator, this))
	}
}

// Checks if can move in <dr, dc> direction
// +ve dr for upwards -ve for downwards
// -ve dc for leftwards, +ve for rightwards
func (this *N2048) can_apply_helper(dr, dc int) bool {

	for row := 0; row < 4; row++ {
		next_row := row + dr
		if next_row < 0 || next_row > 3 {
			continue
		}
		for column := 0; column < 4; column++ {
			next_column := column + dc
			if next_column < 0 || next_column > 3 {
				continue
			}
			current_value := this.board[row][column]
			next_value := this.board[next_row][next_column]
			if current_value != 0 {
				if current_value == next_value || next_value == 0 {
					return true
				}
			}
		}
	}

	return false
}

func (this *N2048) inc_cost(value int) {
	this.path_cost += value
}

func stack_inserter(s Stack, value int) int {
	sum := 0
	if (!s.empty()) && (value == s.peak()) {
		value = value * 2
		sum = value
		s.pop()
	}
	s.push(value)
	return sum
}

func (this *N2048) apply(operator int) Node {
	switch operator {
	case RIGHT:
		return this.apply_helper_horiz(0, 4, 1, operator)
	case LEFT:
		return this.apply_helper_horiz(3, -1, -1, operator)
	case UP:
		return this.apply_helper_vert(3, -1, -1, operator)
	case DOWN:
		return this.apply_helper_vert(0, 4, 1, operator)
	default:
		panic("Apply passed invalid operator value")
	}

}

// Move tiles along vertical axis
// +ve dc for down -ve for up
func (this *N2048) apply_helper_vert(row_start, row_end, dr, operator int) Node {
	child := &N2048{Grid{}, this.max, this, operator,
		this.path_cost, this.depth + 1}
	stack := make_stack(4)
	for column := 0; column != 4; column++ {
		for row := row_start; row != row_end; row += dr {
			value := stack_inserter(stack, this.board[row][column])
			if value > child.max {
				child.max = value // update max if need be
			}

			child.path_cost += value // update operator cost
		}
		// Insert in reverse order
		for reverse := row_end - dr; !stack.empty(); reverse -= dr {
			child.board[reverse][column] = stack.pop()
		}
	}
	add_tile(&child.board) // add a two in the first free corner
	return child

}

// Move tiles along horizontal axis
// +ve dc for right -ve for left
func (this *N2048) apply_helper_horiz(column_start, column_end, dc, operator int) Node {
	child := &N2048{Grid{}, this.max, this, operator,
		this.path_cost, this.depth + 1}
	stack := make_stack(4)
	for row := 0; row != 4; row++ {
		for column := column_start; column != column_end; column += dc {
			value := stack_inserter(stack, this.board[row][column])
			if value > child.max {
				child.max = value // update max if need be
			}

			child.path_cost += value // update operator cost
		}
		// insert in reverse order
		for reverse := column_end - dc; !stack.empty(); reverse -= dc {
			child.board[row][reverse] = stack.pop()
		}
	}
	add_tile(&child.board) // add a two in the first free corner
	return child
}

// adds a 2 at the first free corner if there is one
func add_tile(g *Grid) {
	r, c, ok := first_empty_corner(g)
	if ok {
		g[r][c] = 2
	}
}

// Returns 3-tuple (r, c, ok) where r is row number c is column number
// and ok is true if there is an empty corner
// Corner order : Top-left and then clock-wise
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

// Returns the max value in a grid.
func max_of_grid(g *Grid) int {
	max := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g[i][j] > max {
				max = g[i][j]
			}
		}
	}
	return max
}
