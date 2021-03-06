package main

import (
	"fmt"
	"math"
)

// Grid is an unsigned 64 bit integer which represents
// a packed grid.
// Where each 4 bits represents a power of two, they corrospond to a cell element
// in row major ordering.
type Grid uint64

var gobal_hash map[Grid]bool

// grid_ins inserts in the grid,
// val must be a power of two.
func (grid Grid) grid_ins(row, col, val int) Grid {
	x := row*16 + col*4
	pow := math.Log2(float64(val))
	y := Grid(pow) << Grid(x)
	mask := ^(uint64(0xf) << uint64(x))
	grid = grid&Grid(mask) | Grid(y)
	return grid
}

// grid_access returns the element at row and col
func (grid Grid) grid_access(row, col int) int {
	var x float64 = float64(row*16 + col*4)
	var sum float64 = 0
	var i float64
	for i = 0; i < 4; i++ {
		if (grid & (1 << Grid(x+i))) > 0 {
			sum += math.Pow(2, i)
		}
	}
	if sum == 0 {
		return 0
	}
	return int(math.Pow(2, sum))
}

func (grid Grid) display() [4][4]int {
	var array [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			array[i][j] = grid.grid_access(i, j)
		}
	}
	return array
}

type N2048 struct {
	board      Grid   // The grid
	max        int    // maximum value in board
	parent     *N2048 // nil if root
	operator   int    // operator applied on parent to reach this node
	path_cost  int    // cost from initial state to here
	depth      int    // distance from initial state
	score      int    // heuristic score
	score_flag bool   //score flag is set
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
	translation_map := make(map[int]string)
	translation_map[LEFT] = "left"
	translation_map[RIGHT] = "right"
	translation_map[UP] = "up"
	translation_map[DOWN] = "down"
	translation_map[START] = "START"
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
			current_value := this.board.grid_access(row, column)
			next_value := this.board.grid_access(next_row, next_column)
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
// +ve dr for down -ve for up
func (this *N2048) apply_helper_vert(row_start, row_end, dr, operator int) Node {
	child := &N2048{Grid(0), this.max, this, operator,
		this.path_cost, this.depth + 1, 0, false}
	stack := make_stack(4)
	for column := 0; column != 4; column++ {
		for row := row_start; row != row_end; row += dr {
			value := this.board.grid_access(row, column)
			if value != 0 {
				// ignore zeros
				stack.push(value)
			}
		}
		// Insert in reverse order
		for reverse := row_end - dr; !stack.empty(); {
			current_value := child.board.grid_access(reverse, column)
			if current_value == 0 {
				child.board = child.board.grid_ins(reverse, column, stack.pop())
				continue
			} else if current_value == stack.peak() {
				new_value := stack.pop() * 2
				child.board = child.board.grid_ins(reverse, column, new_value)
				child.path_cost += new_value
				if new_value > child.max {
					child.max = new_value
				}
			}
			reverse -= dr // decrement if not equal or merged only
		}
	}
	add_tile(&child.board) // add a two in the first free corner
	return child

}

// Move tiles along horizontal axis
// +ve dc for right -ve for left
func (this *N2048) apply_helper_horiz(column_start, column_end, dc, operator int) Node {
	child := &N2048{Grid(0), this.max, this, operator,
		this.path_cost, this.depth + 1, 0, false}
	stack := make_stack(4)
	for row := 0; row != 4; row++ {
		for column := column_start; column != column_end; column += dc {
			value := this.board.grid_access(row, column)
			if value != 0 {
				// ignore zeros
				stack.push(value)
			}
		}
		// insert in reverse order
		for reverse := column_end - dc; !stack.empty(); {
			current_value := child.board.grid_access(row, reverse)
			if current_value == 0 {
				child.board = child.board.grid_ins(row, reverse, stack.pop())
				continue
			} else if current_value == stack.peak() {
				new_value := stack.pop() * 2
				child.board = child.board.grid_ins(row, reverse, new_value)
				child.path_cost += new_value
				if new_value > child.max {
					child.max = new_value
				}
			}
			reverse -= dc // decrement only if not equal or merged
		}
	}
	add_tile(&child.board) // add a two in the first free corner
	return child
}

// adds a 2 at the first free corner if there is one
func add_tile(g *Grid) {
	r, c, ok := first_empty_corner(g)
	if ok {
		*g = g.grid_ins(r, c, 2)
	}
}

// Returns 3-tuple (r, c, ok) where r is row number c is column number
// and ok is true if there is an empty corner
// Corner order : Top-left and then clock-wise
func first_empty_corner(g *Grid) (r, c int, ok bool) {
	if g.grid_access(0, 0) == 0 {
		return 0, 0, true
	} else if g.grid_access(0, 3) == 0 {
		return 0, 3, true
	} else if g.grid_access(3, 3) == 0 {
		return 3, 3, true
	} else if g.grid_access(3, 0) == 0 {
		return 3, 0, true
	} else {
		return -1, -1, false
	}
}
