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
	NodeA Node
	NodeB Node
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
	CreateUndirectedNodePair() ([]NodeSet, error)
	CreateUndirectedEdge() ([]NodeSet, error)
}

func (g *NewGraph) CreateUndirectedNodePair() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 {
		for i := range g.NodeA {
			for j := i; j < len(g.NodeA); j++ {
				if j != i {
					// noraml
					nodeSetter.NodeA = g.NodeA[i]
					nodeSetter.NodeB = g.NodeA[j]
					g.NodeSets = append(g.NodeSets, nodeSetter)
					// invert
					nodeSetter.NodeA = g.NodeA[j]
					nodeSetter.NodeB = g.NodeA[i]
					g.NodeSets = append(g.NodeSets, nodeSetter)
				}
			}
		}
	} else {
		return nil, errors.New("CreateUndirectedNodePair: can not create pair")
	}
	return g.NodeSets, nil
}

func (g *NewGraph) CreateUndirectedEdge() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 && len(g.NodeB) > 1 {
		// noramal
		for i := range g.NodeA {
			for j := range g.NodeB {
				nodeSetter.NodeA = g.NodeA[i]
				nodeSetter.NodeB = g.NodeB[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
		// invert
		for i := range g.NodeB {
			for j := range g.NodeA {
				nodeSetter.NodeA = g.NodeB[i]
				nodeSetter.NodeB = g.NodeA[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
	} else {
		return nil, errors.New("CreateUndirectedEdge: can not create edge")
	}
	return g.NodeSets, nil
}

func ShowUndirectedGraphOfNodePair(g UndirectedGraph) {
	if _, err := g.CreateUndirectedNodePair(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

func ShowUndirectedGraphOfEdge(g UndirectedGraph) {
	if _, err := g.CreateUndirectedEdge(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

// endregion UndirectedGraph

// region DirectedGraph
type DirectedGraph interface {
	PrintGraph()
	CreateDirectedNodePair() ([]NodeSet, error)
	CreateDirectedEdge() ([]NodeSet, error)
}

func (g *NewGraph) CreateDirectedNodePair() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 {
		for i := range g.NodeA {
			for j := i; j < len(g.NodeA); j++ {
				if j != i {
					// noraml
					nodeSetter.NodeA = g.NodeA[i]
					nodeSetter.NodeB = g.NodeA[j]
					g.NodeSets = append(g.NodeSets, nodeSetter)
				}
			}
		}
	} else {
		return nil, errors.New("CreateDirectedNodePair: can not create pair")
	}
	return g.NodeSets, nil
}

func (g *NewGraph) CreateDirectedEdge() ([]NodeSet, error) {
	nodeSetter := NodeSet{}
	if len(g.NodeA) > 1 && len(g.NodeB) > 1 {
		// noramal
		for i := range g.NodeA {
			for j := range g.NodeB {
				nodeSetter.NodeA = g.NodeA[i]
				nodeSetter.NodeB = g.NodeB[j]
				g.NodeSets = append(g.NodeSets, nodeSetter)
			}
		}
	} else {
		return nil, errors.New("CreateDirectedEdge: can not create edge")
	}
	return g.NodeSets, nil
}

func ShowDirectedGraphOfNodePair(g DirectedGraph) {
	if _, err := g.CreateDirectedNodePair(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

func ShowDirectedGraphOfEdge(g DirectedGraph) {
	if _, err := g.CreateDirectedEdge(); err != nil {
		log.Fatal(err)
	}
	g.PrintGraph()
}

// endregion DirectedGraph
