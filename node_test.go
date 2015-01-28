package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGridCreation(t *testing.T) {
	grid := GenGrid()
	two_count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if grid.grid_access(i, j) == 2 {
				two_count++
			} else {
				assert.Equal(t, grid.grid_access(i, j), 0, "Every cell value should be either a two or zero")
			}
		}
	}
	assert.Equal(t, two_count, 2, "Exactly two 2 tiles should be present in initial grid")
}

func TestInitialState(t *testing.T) {
	grid := GenGrid()
	problem := P2048{16, &grid}
	state := problem.initial_state()
	assert.Equal(t, state.get_operator(), START, "Initial state should have START value")
	assert.Equal(t, state.get_depth(), 0, "Initial state should have zero depth")

}

func grid_from_array(data *[4][4]int) Grid {
	grid := Grid(0)
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			grid = grid.grid_ins(row, column, data[row][column])
		}
	}
	return grid
}

func all_move() Grid {
	grid_array := [4][4]int{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}
	return grid_from_array(&grid_array)

}

func no_move_grid() Grid {
	grid_array := [4][4]int{
		{2, 4, 8, 16},
		{4, 8, 16, 32},
		{8, 16, 32, 64},
		{128, 265, 1024, 2048}}
	return grid_from_array(&grid_array)

}

func TestCanExpand(t *testing.T) {
	no_expand := N2048{no_move_grid(), 2048, nil, START, 0, 0, 0, false}
	assert.False(t, no_expand.can_apply(LEFT), "blocked grid can't move left")
	assert.False(t, no_expand.can_apply(RIGHT), "blocked grid can't move right")
	assert.False(t, no_expand.can_apply(UP), "blocked grid can't move up")
	assert.False(t, no_expand.can_apply(DOWN), "blocked grid can't move down")

	all_expand := N2048{all_move(), 2, nil, START, 0, 0, 0, false}
	assert.True(t, all_expand.can_apply(LEFT), "free grid should move left")
	assert.True(t, all_expand.can_apply(RIGHT), "free grid should move right")
	assert.True(t, all_expand.can_apply(UP), "free grid should move up")
	assert.True(t, all_expand.can_apply(DOWN), "free grid should move down")
}

func TestMoveRight(t *testing.T) {

	grid_array := [4][4]int{
		{2, 0, 8, 16},
		{4, 0, 0, 32},
		{0, 0, 0, 16},
		{256, 256, 0, 256}}
	node := N2048{grid_from_array(&grid_array),
		265, nil, START, 0, 0, 0, false}
	assert.True(t, node.can_apply(LEFT), "node should be able to move left")
	assert.True(t, node.can_apply(RIGHT), "node should be able to move ")
	assert.True(t, node.can_apply(DOWN), "node should be able to move down")
	assert.True(t, node.can_apply(UP), "node should be able to down")

	resulting_node := node.apply(RIGHT).(*N2048)
	assert.Equal(t, resulting_node.get_parent(), &node, "Parent of child should be node")
	assert.Equal(t, resulting_node.get_depth(), node.depth+1, "Child's depth should be parent + 1")
	assert.Equal(t, resulting_node.get_operator(), RIGHT, "We applied a right operator")

	expected_array := [4][4]int{
		{2, 2, 8, 16},
		{0, 0, 4, 32},
		{0, 0, 0, 16},
		{0, 0, 256, 512}}
	assert.Equal(t, grid_from_array(&expected_array),
		resulting_node.board, "Result array mismatch")
}

func TestMoveLeft(t *testing.T) {
	grid_array := [4][4]int{
		{2, 0, 2, 2},
		{4, 0, 0, 32},
		{0, 0, 0, 16},
		{256, 0, 0, 256}}
	node := N2048{grid_from_array(&grid_array),
		265, nil, START, 0, 0, 0, false}
	assert.True(t, node.can_apply(LEFT), "node should be able to move left")
	assert.True(t, node.can_apply(RIGHT), "node should be able to move ")
	assert.True(t, node.can_apply(DOWN), "node should be able to move down")
	assert.True(t, node.can_apply(UP), "node should be able to down")

	resulting_node := node.apply(LEFT).(*N2048)
	assert.Equal(t, resulting_node.get_parent(), &node, "Parent of child should be node")
	assert.Equal(t, resulting_node.get_depth(), node.depth+1, "Child's depth should be parent + 1")
	assert.Equal(t, resulting_node.get_operator(), LEFT, "We applied a left operator")

	expected_array := [4][4]int{
		{4, 2, 0, 2},
		{4, 32, 0, 0},
		{16, 0, 0, 0},
		{512, 0, 0, 0}}
	assert.Equal(t, grid_from_array(&expected_array),
		resulting_node.board, "Result array mismatch")
}

func TestMoveDown(t *testing.T) {
	grid_array := [4][4]int{
		{4, 8, 2, 2},
		{0, 8, 2, 0},
		{4, 8, 0, 0},
		{4, 8, 0, 0}}
	node := N2048{grid_from_array(&grid_array),
		265, nil, START, 0, 0, 0, false}
	assert.True(t, node.can_apply(LEFT), "node should be able to move left")
	assert.True(t, node.can_apply(RIGHT), "node should be able to move ")
	assert.True(t, node.can_apply(DOWN), "node should be able to move down")
	assert.True(t, node.can_apply(UP), "node should be able to down")

	resulting_node := node.apply(DOWN).(*N2048)
	assert.Equal(t, resulting_node.get_parent(), &node, "Parent of child should be node")
	assert.Equal(t, resulting_node.get_depth(), node.depth+1, "Child's depth should be parent + 1")
	assert.Equal(t, resulting_node.get_operator(), DOWN, "We applied a DOWN operator")

	expected_array := [4][4]int{
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 16, 0, 0},
		{8, 16, 4, 2}}
	assert.Equal(t, grid_from_array(&expected_array),
		resulting_node.board, "Result array mismatch")

}

func TestMoveUp(t *testing.T) {
	grid_array := [4][4]int{
		{4, 0, 2, 32},
		{0, 0, 16, 0},
		{4, 0, 0, 32},
		{4, 8, 256, 0}}
	node := N2048{grid_from_array(&grid_array),
		265, nil, START, 0, 0, 0, false}
	assert.True(t, node.can_apply(LEFT), "node should be able to move left")
	assert.True(t, node.can_apply(RIGHT), "node should be able to move ")
	assert.True(t, node.can_apply(DOWN), "node should be able to move down")
	assert.True(t, node.can_apply(UP), "node should be able to down")

	resulting_node := node.apply(UP).(*N2048)
	assert.Equal(t, resulting_node.get_parent(), &node, "Parent of child should be node")
	assert.Equal(t, resulting_node.get_depth(), node.depth+1, "Child's depth should be parent + 1")
	assert.Equal(t, resulting_node.get_operator(), UP, "We applied an UP operator")

	expected_array := [4][4]int{
		{8, 8, 2, 64},
		{4, 0, 16, 0},
		{0, 0, 256, 0},
		{0, 0, 0, 2}}
	assert.Equal(t, grid_from_array(&expected_array),
		resulting_node.board, "Result array mismatch")
}
