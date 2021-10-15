package graph

func ExamplePrint() {
	sampleNodeA := []Node{
		{ID: 1, Element: "Node-A"},
		{ID: 2, Element: "Node-B"},
		{ID: 3, Element: "Node-C"},
		{ID: 4, Element: "Node-D"},
	}
	sampleNodeB := []Node{
		{ID: 5, Element: "Node-E"},
		{ID: 6, Element: "Node-F"},
	}
	testGraph := &NewGraph{
		NodeA: sampleNodeA,
		NodeB: sampleNodeB,
	}
	testGraph.CreateUndirectedVertexesPair()
	testGraph.CreateUndirectedEdge()
	testGraph.PrintGraph()
	// Output:
	// {{1 Node-A} {2 Node-B}}
	// {{1 Node-A} {3 Node-C}}
	// {{1 Node-A} {4 Node-D}}
	// {{2 Node-B} {3 Node-C}}
	// {{2 Node-B} {4 Node-D}}
	// {{3 Node-C} {4 Node-D}}
	// {{1 Node-A} {5 Node-E}}
	// {{1 Node-A} {6 Node-F}}
	// {{2 Node-B} {5 Node-E}}
	// {{2 Node-B} {6 Node-F}}
	// {{3 Node-C} {5 Node-E}}
	// {{3 Node-C} {6 Node-F}}
	// {{4 Node-D} {5 Node-E}}
	// {{4 Node-D} {6 Node-F}}
}
