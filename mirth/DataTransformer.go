package mirth

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
)

//"fmt"

type Node struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	X     int
	Y     int
}

type Edge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	T      string `json:"type"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges`
}

func ToGraph(channels map[string]Channel) *Graph {
	edges := make([]Edge, 0, len(channels))
	nodes := make([]Node, 0, len(channels))

	for _, v := range channels {

		// one node for every channel
		n := &Node{
			Id:    v.Id,
			Label: v.Name,
		}
		nodes = append(nodes, *n)

		for d := range v.DestinationIds {

			//one edge for every channel destination
			e := &Edge{
				Id:     fmt.Sprintf("%d", len(edges)+10001),
				Source: v.Id,
				Target: v.DestinationIds[d],
				T:      "arrow",
			}
			edges = append(edges, *e)
		}

	}

	graph := new(Graph)
	graph.Edges = edges
	graph.Nodes = nodes

	return graph
}

func ToGraphJson(channels map[string]Channel) string {
	g := ToGraph(channels)
	fmt.Println(g)
	b, err := json.Marshal(g)
	check(err)
	return string(b)
}
