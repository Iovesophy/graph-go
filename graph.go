package graph

import (
	"errors"
	"fmt"
	"log"
)

type NewGraph struct {
	NodeSets []NodeSet
	NodeA    []Node
	NodeB    []Node
}

type Node struct {
	ID      interface{}
	Element interface{}
}

type NodeSet struct {
	VertexA Node
	VertexB Node
}

// region Check Interface
var (
	_ UndirectedGraph = (*NewGraph)(nil)
	_ DirectedGraph   = (*NewGraph)(nil)
)

// endregion Check Interface

// region common

// print any graph
func (g *NewGraph) PrintGraph() {
	for i := range g.NodeSets {
		fmt.Println(g.NodeSets[i])
	}
}

// endregion common

// region UndirectedGraph
type UndirectedGraph interface {
	PrintGraph()
	CreateUndirectedVertexesPair() ([]NodeSet, error)
	CreateUndirectedEdge() ([]NodeSet, error)
}

func (g *NewGraph) CreateUndirectedVertexesPair() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 {
		for i := range g.NodeA {
			for j := i; j < len(g.NodeA); j++ {
				if j != i {
					nodeSetter.VertexA = g.NodeA[i]
					nodeSetter.VertexB = g.NodeA[j]
					g.NodeSets = append(g.NodeSets, nodeSetter)
				}
			}
		}
	} else {
		return nil, errors.New("CreateUndirectedVertexesPair: can not create pair")
	}
	return g.NodeSets, nil
}

func (g *NewGraph) CreateUndirectedEdge() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 && len(g.NodeB) > 1 {
		for i := range g.NodeA {
			for j := range g.NodeB {
				nodeSetter.VertexA = g.NodeA[i]
				nodeSetter.VertexB = g.NodeB[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
	} else {
		return nil, errors.New("CreateUndirectedEdge: can not create edge")
	}
	return g.NodeSets, nil
}

func ShowUndirectedGraph(g UndirectedGraph) {
	if _, err := g.CreateUndirectedVertexesPair(); err != nil {
		log.Fatal(err)
	}
	if _, err := g.CreateUndirectedEdge(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

// endregion UndirectedGraph

// region DirectedGraph
type DirectedGraph interface {
	PrintGraph()
	CreateDirectedVertexesPair() ([]NodeSet, error)
	CreateDirectedEdge() ([]NodeSet, error)
}

func (g *NewGraph) CreateDirectedVertexesPair() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 {
		for i := range g.NodeA {
			for j := i; j < len(g.NodeA); j++ {
				if j != i {
					nodeSetter.VertexA = g.NodeA[i]
					nodeSetter.VertexB = g.NodeA[j]
					g.NodeSets = append(g.NodeSets, nodeSetter)
				}
			}
		}
	} else {
		return nil, errors.New("CreateDirectedVertexesPair: can not create pair")
	}
	return g.NodeSets, nil
}

func (g *NewGraph) CreateDirectedEdge() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 && len(g.NodeB) > 1 {
		for i := range g.NodeA {
			for j := range g.NodeB {
				nodeSetter.VertexA = g.NodeA[i]
				nodeSetter.VertexB = g.NodeB[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
	} else {
		return nil, errors.New("CreateDirectedEdge: can not create edge")
	}
	return g.NodeSets, nil
}

func ShowDirectedGraph(g DirectedGraph) {
	if _, err := g.CreateDirectedVertexesPair(); err != nil {
		log.Fatal(err)
	}
	if _, err := g.CreateDirectedEdge(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

// endregion DirectedGraph
