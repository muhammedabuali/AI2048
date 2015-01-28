package main

func enqueue_at_end(nodes *[]Node, children []Node) []Node {
	*nodes = (*nodes)[1:]
	return append(*nodes, children...)
}

func enqueue_at_front(nodes *[]Node, children []Node) []Node {
	*nodes = (*nodes)[1:]
	return append(children, *nodes...)
}

func depth_limited_search(limit uint64) Strategy {
	return func(nodes *[]Node, children []Node) []Node {
		*nodes = (*nodes)[1:]
		if uint64(children[0].get_depth()) > limit {
			return *nodes
		} else {
			return append(children, *nodes...)
		}
	}
}

func best_fit_enqueue(h Heuristic) Strategy {
	return func(nodes *[]Node, children []Node) []Node {
		heap_down(nodes, h)
		for i := 0; i < len(children); i++ {
			heap_up(nodes, children[i], h)
		}
		return *nodes
	}
}

type ByEval struct {
	nodes []Node
	g     Heuristic
}

func (a ByEval) Len() int           { return len(a.nodes) }
func (a ByEval) Swap(i, j int)      { a.nodes[i], a.nodes[j] = a.nodes[j], a.nodes[i] }
func (a ByEval) Less(i, j int) bool { return a.g(a.nodes[i]) < a.g(a.nodes[j]) }

func heap_up(nodes *[]Node, node Node, score Heuristic) {
	cur := len(*nodes)
	*nodes = append(*nodes, node)
	for cur > 0 {
		if score((*nodes)[(cur-1)/2]) >= score((*nodes)[cur]) {
			(*nodes)[cur] = (*nodes)[cur-1/2]
			(*nodes)[cur-1/2] = node
			cur = (cur - 1) / 2
		} else {
			break
		}
	}
}

func heap_down(nodes *[]Node, score Heuristic) {
	(*nodes)[0] = (*nodes)[len(*nodes)-1]
	*nodes = (*nodes)[:len(*nodes)-1]
	cur := 0
	if len(*nodes) == 0 {
		return
	}
	cur_node := (*nodes)[0]
	for {
		if cur*2+1 >= len(*nodes) {
			break
		}
		if cur*2+2 >= len(*nodes) {
			if score(cur_node) > score((*nodes)[cur*2+1]) {
				(*nodes)[cur] = (*nodes)[cur*2+1]
				(*nodes)[cur*2+1] = cur_node
				break
			} else {
				break
			}
		}
		if score(cur_node) <= score((*nodes)[cur*2+1]) && score(cur_node) <= score((*nodes)[cur*2+2]) {
			break
		}
		if score((*nodes)[cur*2+1]) <= score(cur_node) && score((*nodes)[cur*2+1]) <= score((*nodes)[cur*2+2]) {
			(*nodes)[cur] = (*nodes)[cur*2+1]
			(*nodes)[cur*2+1] = cur_node
			cur = cur*2 + 1
		} else {
			(*nodes)[cur] = (*nodes)[cur*2+2]
			(*nodes)[cur*2+2] = cur_node
			cur = cur*2 + 2
		}
	}
}
