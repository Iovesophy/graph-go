package graph

import "fmt"

var (
	_ Graph = (*Glink)(nil)
)

type Graph interface {
	CreateVertexesPair() []NodeSet
	CreateEdge() []NodeSet
	PrintGraph()
}

type Glink struct {
	NodeSets []NodeSet
	NodeData []Node
	BaseData []Node
}

type Node struct {
	ID      int
	Element string
}

type NodeSet struct {
	Vertex1 Node
	Vertex2 Node
}

func (g *Glink) CreateVertexesPair() []NodeSet {
	nodeSetter := NodeSet{}
	for i := range g.NodeData {
		for j := i; j < len(g.NodeData); j++ {
			if j != i {
				nodeSetter.Vertex1 = g.NodeData[i]
				nodeSetter.Vertex2 = g.NodeData[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
	}
	return g.NodeSets
}

func (g *Glink) CreateEdge() []NodeSet {
	nodeSetter := NodeSet{}
	for i := range g.NodeData {
		for j := range g.BaseData {
			nodeSetter.Vertex1 = g.NodeData[i]
			nodeSetter.Vertex2 = g.BaseData[j]
			g.NodeSets = append(g.NodeSets, nodeSetter)
		}
	}
	return g.NodeSets
}

func (g *Glink) PrintGraph() {
	for i := range g.NodeSets {
		fmt.Println(g.NodeSets[i])
	}
}

func DoGraph(g Graph) {
	_ = g.CreateVertexesPair()
	_ = g.CreateEdge()
	g.PrintGraph()
}
