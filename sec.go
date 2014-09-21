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
}
