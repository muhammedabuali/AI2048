package main

import (
	"math/rand"
)

type P2048 struct {
}

type N2048 struct {
	board [4][4]int
}

func (*P2048) init_state() N2048 {
	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	var node N2048 = N2048{board}
	return node
}
