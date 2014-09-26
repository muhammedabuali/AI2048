package AI2048

import "sort"

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
		combined := append(nodes, children...)
		eval := func(n Node) int {
			return n.get_path_cost() + h(n)
		}
		sort.Sort(ByEval{combined, eval})
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
