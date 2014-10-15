package main

import (
	//"fmt"
	"math"
)

func greedy_heuristic_1(n Node) int {
	return estimated_cost(n.(*N2048), global_problem.(*P2048))
}

func greedy_heuristic_2(n Node) int {
	return 1
}

func astar_heuristic_1(n Node) int {
	return n.get_path_cost() + estimated_cost(n.(*N2048), global_problem.(*P2048))
}

func astar_heuristic_2(n Node) int {
	return n.get_path_cost() + estimated_cost2(n.(*N2048), global_problem.(*P2048))
}

/* estimate coast to reach goal by difference between
minimal overall score and current essential socre
essential score: non redundant merges
*/
func estimated_cost(node *N2048, proplem *P2048) int {
	if node.score_flag {
		return node.score
	}
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
	node.score = estimated
	node.score_flag = true
	return int(estimated)
}

/* adds another term to the first heuristic that detects if redundant
moves are inevitable
*/
func estimated_cost2(node *N2048, proplem *P2048) int {
	if node.score_flag {
		return node.score
	}
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
	var i int

	for i = target; i > 1; i-- {
		// only count useful meges
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
	//check redundant moves
	check := false
	for l := i - 1; l > 0; l-- {
		if counts[l] >= 3 {
			//alot of redundant numbers => check redundant merges
			check = true
		}
	}
	redundant_score := 0
	if check {
		last_level := i - 1
		last_num := int(math.Pow(2, float64(last_level)))
		column := 0
		cur := 0
		/* if you have redundant merges in all directions
		add score to this node */

		for j := 0; j < 4; j++ { //columns
			cur = current_board[0][j]
			for k := 1; k < 3; k++ { //rows
				if current_board[k][j] != cur || current_board[k][j] == 0 {
					cur = current_board[k][j]
				} else { //merge found
					if cur < last_num { // redundant merge
						column = cur * 2 //store redundant score
						break
					}
				}
			}
		}
		for j := 0; j < 4; j++ { //rows
			cur = current_board[j][0]
			for k := 1; k < 3; k++ { //columns
				if current_board[j][k] != cur || current_board[j][k] == 0 {
					cur = current_board[j][k]
				} else { //merge found
					if cur < last_num { // redundant merge
						row := cur * 2 //check rows
						if row > column {
							redundant_score = column
						} else {
							redundant_score = row
						}
					}
				}
			}
		}
	}
	estimated := total_cost - current_cost + redundant_score
	node.score = estimated
	node.score_flag = true
	return int(estimated)
}
