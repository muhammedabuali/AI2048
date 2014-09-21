package main

func bfs(nodes []Node, children []Node) []Node {
	return append(nodes, children...)
}

func bfs(nodes []Node, children []Node) []Node {
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
