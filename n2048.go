package AI2048

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
	path := make(Path, node.depth)
	curr_node := node
	translation_map := make(map[int]string)
	translation_map[LEFT] = "left"
	translation_map[RIGHT] = "right"
	translation_map[UP] = "up"
	translation_map[DOWN] = "down"
	translation_map[START] = "START"
	for i := node.depth - 1; i >= 0; i-- {
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
			"unknown argument passed '%v' to can_apply on N2048:\n%v  ",
			operator, this))
	}
}

// Checks if can move in <dr, dc> direction
// +ve dr for upwards -ve for downwards
// -ve dc for leftwards, +ve for rightwards
func (this *N2048) can_apply_helper(dr, dc int) bool {

	for row := 0; row < 4; row++ {
		next_row := row + dr
		if next_row < 0 || next_row > 4 {
			continue
		}
		for column := 0; column < 4; column++ {
			next_column := column + dc
			if next_column < 0 || next_column > 4 {
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

func (this *N2048) apply(operator int) Node {
	child := &N2048{this.board, this.max, this, operator,
		this.path_cost, this.depth + 1}

	return child
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
