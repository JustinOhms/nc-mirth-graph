package mirth

import (
	//"encoding/json"
	"fmt"
)

//"fmt"

type node struct {
	id    string
	label string
	x     int
	y     int
}

type edge struct {
	id     string
	source string
	target string
	t      string
}

type Graph struct {
	Nodes []node
	Edges []edge
}

func ToGraphJson(channels map[string]Channel) *Graph {
	edges := make([]edge, 0, len(channels))
	nodes := make([]node, 0, len(channels))

	for _, v := range channels {

		// one node for every channel
		n := &node{
			id:    v.Id,
			label: v.Name,
		}
		nodes = append(nodes, *n)

		for d := range v.DestinationIds {

			//one edge for every channel destination
			e := &edge{
				id:     fmt.Sprintf("%d", len(edges)+10001),
				source: v.Id,
				target: v.DestinationIds[d],
				t:      "arrow",
			}
			edges = append(edges, *e)
		}

	}

	graph := &Graph{
		Nodes: nodes,
		Edges: edges,
	}
	//	fmt.Println(nodes)
	//	fmt.Println(edges)
	//	fmt.Println(graph)
	return graph
}
