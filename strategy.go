package main

import "sort"

func bfs(nodes []Node, children []Node) []Node {
	return append(nodes, children...)
}

func dfs(nodes []Node, children []Node) []Node {
	return append(children, nodes...)
}

func depth_limited_search(limit int) Strategy {
	return func(nodes []Node, children []Node) []Node {
		if children[0].get_depth() > limit {
			return nodes
		} else {
			return append(children, nodes...)
		}
	}
}

type evaluation_func func(Node) int

type ByEval struct {
	nodes []Node
	g     evaluation_func
}

func (a ByEval) Len() int           { return len(a.nodes) }
func (a ByEval) Swap(i, j int)      { a.nodes[i], a.nodes[j] = a.nodes[j], a.nodes[i] }
func (a ByEval) Less(i, j int) bool { return a.g(a.nodes[i]) < a.g(a.nodes[j]) }

func greedy_search(h evaluation_func) Strategy {
	return func(nodes []Node, children []Node) []Node {
		combined := append(nodes, children...)
		sort.Sort(ByEval{combined, h})
		return combined
	}
}

func a_star(h evaluation_func) Strategy {
	return func(nodes []Node, children []Node) []Node {
		combined := append(nodes, children...)
		eval := func(n Node) int {
			return n.get_path_cost() + h(n)
		}
		sort.Sort(ByEval{combined, eval})
		return combined
	}
}
