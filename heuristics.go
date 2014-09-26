package main

func greedy_heuristic_1(n Node) int {
	return 1
}

func greedy_heuristic_2(n Node) int {
	return 1
}

func astar_heuristic_1(n Node) int {
	return n.get_path_cost() + 1
}

func astar_heuristic_2(n Node) int {
	return n.get_path_cost() + 1
}
