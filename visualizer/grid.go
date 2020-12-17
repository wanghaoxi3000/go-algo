/*Package visulizer
vscode 可视化调试插件 debug visulizer 数据结构定义
Gird 视图
*/

package visulizer

import (
	"encoding/json"
	"strconv"
)

func NewGrid() *Grid {
	return &Grid{
		Kind:    map[string]bool{"grid": true},
		Rows:    make([]*GridRow, 0),
		Markers: make([]*Marker, 0),
	}
}

func NewGridFromSliceInt(s []int, marks ...int) *Grid {
	grid := NewGrid()

	markmap := make(map[int]bool, len(marks))
	for _, v := range marks {
		markmap[v] = true
	}

	node := make([]*GridNode, len(s))
	for i, v := range s {
		node[i] = &GridNode{
			Tag: strconv.Itoa(v),
		}
		if markmap[v] {
			grid.AddMarker("Target", 0, i)
			markmap[v] = false
		}
	}

	row := &GridRow{Columns: node}

	grid.Rows = append(grid.Rows, row)
	return grid
}

//Grid 网格列表视图
type Grid struct {
	Kind    map[string]bool `json:"kind,omitempty"`
	Rows    []*GridRow      `json:"rows,omitempty"`
	Markers []*Marker       `json:"markers,omitempty"`
}

type GridRow struct {
	Label   string      `json:"label,omitempty"`
	Columns []*GridNode `json:"columns,omitempty"`
}

type GridNode struct {
	Tag     string `json:"tag,omitempty"`
	Content string `json:"content,omitempty"`
	Color   string `json:"color,omitempty"`
}

type Marker struct {
	ID     string `json:"id,omitempty"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
	Label  string `json:"label,omitempty"`
	Color  string `json:"color,omitempty"`
}

func (g *Grid) Visulize() string {
	rs, _ := json.Marshal(g)
	return string(rs)
}

func (g *Grid) AddMarker(id string, row, col int) *Marker {
	marker := &Marker{
		ID:     id,
		Row:    row,
		Column: col,
	}

	g.Markers = append(g.Markers, marker)
	return marker
}
