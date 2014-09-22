package main

import "fmt"

func hello() {
	fmt.Println("hello")
	var a []int
	fmt.Println(len(a))
	a = append(a, 4)
	fmt.Println(len(a))
	a = make([]int, 10)
	fmt.Println(len(a))
	a = a[10:]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	var n Node
	fmt.Println(n == nil)
	type Hey []int

	// node experiment
	board := [4][4]int{
		{2, 0, 2, 0},
		{2, 2, 2, 2},
		{2, 0, 0, 0},
		{2, 4, 0, 2}}
	var root *N2048
	var node N2048 = N2048{board, 0, root, 0, 0, 0}
	var node1 N2048 = make_node(node, 1)
	var node5 N2048 = make_node(node1, 1)
	for i := 0; i < 4; i++ {
		fmt.Println(node.board[i])
	}
	fmt.Println("*******")
	for i := 0; i < 4; i++ {
		fmt.Println(node1.board[i])
	}
	fmt.Println("*******")
	for i := 0; i < 4; i++ {
		fmt.Println(node5.board[i])
	}
	var p Problem
	p = &P2048{8}
	var sol Node = general_search(p, bfs)
	var solution = sol.(N2048)
	var parent *N2048 = solution.parent
	fmt.Println("***sol***")
	for i := 0; i < 4; i++ {
		fmt.Println(solution.board[i])
	}
	for parent != nil {
		fmt.Println("***sol***")
		for i := 0; i < 4; i++ {
			fmt.Println(parent.board[i])
		}
		parent = parent.parent
	}
}
