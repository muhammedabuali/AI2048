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
	var node N2048 = N2048{board, 0}
	var node1 N2048 = make_node(node, 1)
	var node2 N2048 = make_node(node, 2)
	for i := 0; i < 4; i++ {
		fmt.Println(node.board[i])
	}
	fmt.Println("*******")
	for i := 0; i < 4; i++ {
		fmt.Println(node1.board[i])
	}
	fmt.Println("*******")
	for i := 0; i < 4; i++ {
		fmt.Println(node2.board[i])
	}
}
