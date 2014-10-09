package main

import (
	"fmt"
	"math"
)

func greedy_heuristic_1(n Node) int {
	return 1
}

func greedy_heuristic_2(n Node) int {
	return 1
}

func astar_heuristic_1(n Node) int {
	return estimated_cost(n.(*N2048), global_problem.(*P2048))
}

func astar_heuristic_2(n Node) int {
	return 1
}

func estimated_cost(node *N2048, proplem *P2048) int {
	total_cost := int(float64(proplem.goal) * (math.Log2(float64(proplem.goal)) - 1))
	current_board := node.board.display()
	target := int(math.Log2(float64(proplem.goal)))

	counts := make([]int, target+1, target+1)

	// count occurances of every power
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if current_board[i][j] > 0 {
				power := int(math.Log2(float64(current_board[i][j])))
				counts[power]++
			}
		}
	}
	current_cost := 0
	score := 0
	needed := 1

	for i := target; i > 1; i-- {
		// only count useful merges
		if counts[i] > needed {
			counts[i] = needed
		}
		score = counts[i] * int(math.Pow(2, float64(i))) * (i - 1)
		current_cost += score
		if counts[i] == needed {
			break
		}
		needed = (needed - counts[i]) * 2
	}
	estimated := total_cost - current_cost
	//fmt.Println("hello")
	//fmt.Println("node", node.board.display(), "total cost", total_cost, "current_cost", current_cost)
	return int(estimated)
}
