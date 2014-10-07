package main

import (
	"fmt"
	"math"
	"math/rand"
)

type G2 uint64

var hash map[G2]bool

func (grid G2) grid_ins(row, col, val int) G2 {
	x := row*16 + col*4
	pow := math.Log2(float64(val))
	y := G2(pow) << G2(x)
	mask := ^(0xff << uint(x))
	grid = grid&G2(mask) | G2(y)
	return grid
}

func (grid G2) grid_access(row, col int) int {
	var x float64 = float64(row*16 + col*4)
	var sum float64 = 0
	var i float64
	for i = 0; i < 4; i++ {
		if (grid & (1 << G2(x+i))) > 0 {
			sum += math.Pow(2, i)
		}
	}
	if sum == 0 {
		return 0
	}
	return int(math.Pow(2, sum))
}

func (grid G2) display() [4][4]int {
	var array [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			array[i][j] = grid.grid_access(i, j)
		}
	}
	return array
}

type N2 struct {
	board     G2  // The grid
	max       int // maximum value in board
	parent    *N2 // nil if root
	operator  int // operator applied on parent to reach this node
	path_cost int // cost from initial state to here
	depth     int // distance from initial state
}

func (node *N2) get_parent() Node {
	return node.parent
}

func (node *N2) get_operator() int {
	return node.operator
}

func (node *N2) get_path_cost() int {
	return node.path_cost
}

func (node *N2) get_depth() int {
	return node.depth
}

func (node *N2) get_path() Path {

	path := make(Path, node.depth+1)
	curr_node := node
	translation_map := make(map[int]string)
	translation_map[LEFT] = "left"
	translation_map[RIGHT] = "right"
	translation_map[UP] = "up"
	translation_map[DOWN] = "down"
	translation_map[START] = "START"
	for i := node.depth; i >= 0; i-- {
		fmt.Println(curr_node.board.display())
		fmt.Println(curr_node.get_path_cost())
		path[i] = translation_map[curr_node.operator]
		curr_node = curr_node.parent
	}
	return path
}

func (this *N2) can_apply(operator int) bool {

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
			"unknown argument passed '%v' to can_apply on N2:\n%v",
			operator, this))
	}
}

// Checks if can move in <dr, dc> direction
// +ve dr for upwards -ve for downwards
// -ve dc for leftwards, +ve for rightwards
func (this *N2) can_apply_helper(dr, dc int) bool {

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
			next_value := this.board.grid_access(row, column)
			if current_value != 0 {
				if current_value == next_value || next_value == 0 {
					return true
				}
			}
		}
	}

	return false
}

func (this *N2) inc_cost(value int) {
	this.path_cost += value
}

func stack_inserter2(s Stack, value int, can_merge bool) (int, bool) {
	if value == 0 {
		return 0, can_merge
	}
	sum := 0
	if !can_merge {
		can_merge = true
	} else if (!s.empty()) && (value == s.peak()) {

		value = value * 2
		sum = value
		s.pop()
		can_merge = false
	}
	s.push(value)
	return sum, can_merge
}

func (this *N2) apply(operator int) Node {
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
func (this *N2) apply_helper_vert(row_start, row_end, dr, operator int) Node {
	child := &N2{0, this.max, this, operator,
		this.path_cost, this.depth + 1}
	stack := make_stack(4)
	value := 0
	can_merge := true
	for column := 0; column != 4; column++ {
		for row := row_start; row != row_end; row += dr {
			value, can_merge = stack_inserter2(stack, this.board.grid_access(row, column), can_merge)
			if value > child.max {
				child.max = value // update max if need be
			}

			child.path_cost += value // update operator cost
		}
		// Insert in reverse order
		for reverse := row_end - dr; !stack.empty(); reverse -= dr {
			child.board = child.board.grid_ins(reverse, column, stack.pop())
		}
	}
	add_tile2(&child.board) // add a two in the first free corner
	return child

}

// Move tiles along horizontal axis
// +ve dc for right -ve for left
func (this *N2) apply_helper_horiz(column_start, column_end, dc, operator int) Node {
	child := &N2{G2(0), this.max, this, operator,
		this.path_cost, this.depth + 1}
	stack := make_stack(4)
	value := 0
	can_merge := true
	for row := 0; row != 4; row++ {
		for column := column_start; column != column_end; column += dc {
			value, can_merge = stack_inserter(stack, this.board.grid_access(row, column), can_merge)
			if value > child.max {
				child.max = value // update max if need be
			}

			child.path_cost += value // update operator cost
		}
		// insert in reverse order
		for reverse := column_end - dc; !stack.empty(); reverse -= dc {
			child.board = child.board.grid_ins(row, reverse, stack.pop())
		}
	}
	add_tile2(&child.board) // add a two in the first free corner
	return child
}

// adds a 2 at the first free corner if there is one
func add_tile2(g *G2) {
	r, c, ok := first_empty_corner2(g)
	if ok {
		*g = g.grid_ins(r, c, 2)
	}
}

// Returns 3-tuple (r, c, ok) where r is row number c is column number
// and ok is true if there is an empty corner
// Corner order : Top-left and then clock-wise
func first_empty_corner2(g *G2) (r, c int, ok bool) {
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

// Returns the max value in a grid.
func max_of_grid2(g *G2) int {
	max := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.grid_access(i, j) > max {
				max = g.grid_access(i, j)
			}
		}
	}
	return max
}

/**********           p2048   **********************/
type P2 struct {
	goal int
	grid *G2
}

// Returns the initial state of the board
// As in the Root node of the search tree.
// seed is the seed used by the PRNG.
func (p *P2) initial_state() Node {
	return &N2{*p.grid, 2, nil, START, 0, 0}
}

// Returns true if we have reached M.
func (p *P2) goal_test(n Node) bool {
	return p.goal == n.(*N2).max
}

// Returns all children of a Node
func (*P2) expand(node Node) []Node {
	nodes := make([]Node, 0, 4)
	// Apply all possible operators
	if node.can_apply(LEFT) {
		new_node := node.apply(LEFT)
		fmt.Println(hash)
		fmt.Println((node.(*N2)).board.display())
		fmt.Println((new_node.(*N2)).board.display())
		if !hash[(new_node.(*N2)).board] {
			nodes = append(nodes, new_node)
		}
	}
	if node.can_apply(RIGHT) {
		new_node := node.apply(RIGHT)
		fmt.Println((new_node.(*N2)).board.display())
		if !hash[(new_node.(*N2)).board] {
			nodes = append(nodes, new_node)
		}
	}
	if node.can_apply(DOWN) {
		new_node := node.apply(DOWN)
		fmt.Println((new_node.(*N2)).board.display())
		if !hash[(new_node.(*N2)).board] {
			nodes = append(nodes, new_node)
		}
	}
	if node.can_apply(UP) {
		new_node := node.apply(UP)
		fmt.Println((new_node.(*N2)).board.display())
		if !hash[(new_node.(*N2)).board] {
			nodes = append(nodes, new_node)
		}
	}

	return nodes
}

func Search2(grid *G2, M int, strategy int, visualize bool) (p Path, cost int, nodes uint64) {
	problem := P2{M, grid}
	var (
		target         Node
		success        bool
		nodes_expanded uint64
		goal_path      Path
		path_cost      int
	)
	if strategy == ID {
		target, success, nodes_expanded = iterative_deepening_search(&problem)

	} else {
		quing_func := get_quing_func(strategy)
		target, success, nodes_expanded = general_search(&problem, quing_func)
	}

	if success {
		// reached goal state
		goal_path, path_cost = target.get_path(), target.get_path_cost()

	} else {
		// Failed to reach goal state
		goal_path, path_cost = Path{}, 0
	}

	if visualize {
		//TODO : Implement
	}

	return goal_path, path_cost, nodes_expanded
}

func GenGrid2() G2 {
	var grid G2 = G2(0)
	hash = make(map[G2]bool)
	//rand.Seed(time.Now().UTC().Unix())
	rand.Seed(42)
	r1, c1, r2, c2 := rand.Intn(4), rand.Intn(4), rand.Intn(4), rand.Intn(4)

	for (r1 == r2) && (c2 == c1) {
		c1 = rand.Intn(4)
	}
	grid = grid.grid_ins(r1, c1, 2)
	grid = grid.grid_ins(r2, c2, 2)
	return grid
}
