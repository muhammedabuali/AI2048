package main

import (
	"fmt"
	"math/rand"
)

const (
	BF  = iota
	DF  = iota
	ID  = iota
	GR1 = iota
	GR2 = iota
	AS1 = iota
	AS2 = iota
)

func main() {
	// Generate grid
	grid := GenGrid()
	goal_path, path_cost, nodes_expanded := Search(grid, 4, AS2, true)
	fmt.Printf("Path: %v\nCost: %v\nTotal Nodes Expanded in search: %v\n",
		goal_path, path_cost, nodes_expanded)
}

/*
 * Uses a search strategy to formulate a plan to the goal
 * grid - Initial grid to use
 * M - Power of two that we are trying to reach
 * strategy - one of {BF, DF, ID, GR1, GR2, AS1, AS2} representing the search strategy
 * visualize - true if you want a visual representation
 * return 3-tuple (Path, cost, number of nodes expanded)
 */
func Search(grid Grid, M uint, strategy int, visualize bool) (Path, int, uint64) {
	problem := P2048{1 << M, grid}
	global_problem = &problem
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
		display(goal_path.encode())
	}

	return goal_path, path_cost, nodes_expanded
}

// Returns the quing func for the requested symbol
func get_quing_func(symbol int) Strategy {
	switch symbol {
	case BF:
		return enqueue_at_end
	case DF:
		return enqueue_at_front
	case GR1:
		return best_fit_enqueue(greedy_heuristic_1)
	case GR2:
		return best_fit_enqueue(greedy_heuristic_2)
	case AS1:
		return best_fit_enqueue(astar_heuristic_1)
	default:
		// AS2
		return best_fit_enqueue(astar_heuristic_2)
	}
}

func GenGrid() Grid {
	var grid Grid = Grid(0)
	gobal_hash = make(map[Grid]bool)
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
