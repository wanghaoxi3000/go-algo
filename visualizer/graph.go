/*Package visulizer
vscode 可视化调试插件 debug visulizer 数据结构定义
*/

package visulizer

import (
	"encoding/json"
	"strconv"
)

type Graph struct {
	Kind  map[string]bool  `json:"kind"`
	Nodes []*NodeGraphData `json:"nodes"`
	Edges []*EdgeGraphData `json:"edges"`
}

type GraphNode struct {
	Value int

	Node NodeGraphData
	Edge EdgeGraphData
}

type NodeGraphData struct {
	ID    string `json:"id"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
	Shape string `json:"shape,omitempty"`
}

type EdgeGraphData struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Label  string `json:"label,omitempty"`
	ID     string `json:"id"`
	Color  string `json:"color,omitempty"`
	Dashes bool   `json:"dashes,omitempty"`
}

func NewGraph() *Graph {
	return &Graph{
		Kind:  map[string]bool{"graph": true},
		Nodes: make([]*NodeGraphData, 0),
		Edges: make([]*EdgeGraphData, 0),
	}
}

func NewGraphFromSliceInt(s []int) *Graph {
	graph := NewGraph()

	last := 0
	for i, v := range s {

		id := strconv.Itoa(i)
		label := strconv.Itoa(v)
		graph.Nodes = append(graph.Nodes, &NodeGraphData{
			ID:    id,
			Label: label,
		})

		if i > 0 {
			graph.Edges = append(graph.Edges, &EdgeGraphData{
				From: strconv.Itoa(last),
				To:   id,
			})
		}

		last = i
	}

	return graph
}

func (g *Graph) AddNode(data *GraphNode) {
	g.Nodes = append(g.Nodes, &data.Node)
	g.Edges = append(g.Edges, &data.Edge)
}

func (g *Graph) ToString() string {
	rs, _ := json.Marshal(g)
	return string(rs)
}
