package graph

func ExamplePrintGraph() {
	sampleNode := []Node{
		Node{ID: 1, Element: "Node-A"},
		Node{ID: 2, Element: "Node-B"},
		Node{ID: 3, Element: "Node-C"},
		Node{ID: 4, Element: "Node-D"},
	}
	sampleBase := []Node{
		Node{ID: 5, Element: "base-A"},
		Node{ID: 6, Element: "base-B"},
	}
	f := &Glink{
		NodeData: sampleNode,
		BaseData: sampleBase,
	}
	f.CreateVertexesPair()
	f.CreateEdge()
	f.PrintGraph()
	// Output:
	// {{1 Node-A} {2 Node-B}}
	// {{1 Node-A} {3 Node-C}}
	// {{1 Node-A} {4 Node-D}}
	// {{2 Node-B} {3 Node-C}}
	// {{2 Node-B} {4 Node-D}}
	// {{3 Node-C} {4 Node-D}}
	// {{1 Node-A} {5 base-A}}
	// {{1 Node-A} {6 base-B}}
	// {{2 Node-B} {5 base-A}}
	// {{2 Node-B} {6 base-B}}
	// {{3 Node-C} {5 base-A}}
	// {{3 Node-C} {6 base-B}}
	// {{4 Node-D} {5 base-A}}
	// {{4 Node-D} {6 base-B}}
}
