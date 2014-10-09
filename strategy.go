package main

import (
	"fmt"
	"sort"
)

func enqueue_at_end(nodes []Node, children []Node) []Node {
	return append(nodes, children...)
}

func enqueue_at_front(nodes []Node, children []Node) []Node {
	return append(children, nodes...)
}

func depth_limited_search(limit uint64) Strategy {
	return func(nodes []Node, children []Node) []Node {
		if uint64(children[0].get_depth()) > limit {
			return nodes
		} else {
			return append(children, nodes...)
		}
	}
}

func greedy_enqueue(h Heuristic) Strategy {
	return func(nodes []Node, children []Node) []Node {
		combined := append(nodes, children...)
		sort.Sort(ByEval{combined, h})
		return combined
	}
}

func a_star_enqueue(h Heuristic) Strategy {
	return func(nodes []Node, children []Node) []Node {
		combined := append(children, nodes...)
		eval := func(n Node) int {
			return n.get_path_cost() + h(n)
		}
		/*fmt.Println("before sorting", len(combined))
		fmt.Println(combined)
		for i := 0; i < len(combined); i++ {
			fmt.Println(eval(combined[i]), (combined[i].(*N2048)).board.display())
			fmt.Println(h(combined[i]))
		}
		fmt.Println()*/
		sort.Sort(ByEval{combined, eval})
		/*fmt.Println("after sorting")
		fmt.Println(combined)
		fmt.Println(eval(combined[0]))*/
		return combined
	}
}

type ByEval struct {
	nodes []Node
	g     Heuristic
}

func (a ByEval) Len() int           { return len(a.nodes) }
func (a ByEval) Swap(i, j int)      { a.nodes[i], a.nodes[j] = a.nodes[j], a.nodes[i] }
func (a ByEval) Less(i, j int) bool { return a.g(a.nodes[i]) < a.g(a.nodes[j]) }
