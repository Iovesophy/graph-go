# <img src="https://user-images.githubusercontent.com/15680172/137610607-39e4b8b0-a343-4b22-b37c-0183bbf86dc4.png" data-canonical-src="https://user-images.githubusercontent.com/15680172/137610607-39e4b8b0-a343-4b22-b37c-0183bbf86dc4.png" width="30" height="30" /> Graph-go

[![Go Reference](https://pkg.go.dev/badge/github.com/Iovesophy/graph-go.svg)](https://pkg.go.dev/github.com/Iovesophy/graph-go) [![test](https://github.com/Iovesophy/graph-go/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/Iovesophy/graph-go/actions/workflows/test.yml) [![codecov](https://codecov.io/gh/Iovesophy/graph-go/branch/master/graph/badge.svg?token=LMRN408RMC)](https://codecov.io/gh/Iovesophy/graph-go) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/Iovesophy/graph-go/master/LICENSE)

## sample

```Go
package main

import (
	"fmt"
	"github.com/Iovesophy/graph-go"
)

func main() {
	samplePlace := []graph.Node{
		{ID: 1, Element: "place-A"},
		{ID: 2, Element: "place-B"},
		{ID: 3, Element: "place-C"},
		{ID: 4, Element: "place-D"},
	}
	sampleStation := []graph.Node{
		{ID: 5, Element: "station-A"},
		{ID: 6, Element: "station-B"},
	}
	g := &graph.NewGraph{
		NodeA: samplePlace,
		NodeB: sampleStation,
	}

	val1 , _ := g.CreateUndirectedNodePair()
	val2 , _ := g.CreateUndirectedEdge()
	val3 , _ := g.CreateDirectedNodePair()
	val4 , _ := g.CreateDirectedEdge()
	fmt.Println(val1,val2,val3,val4)

	graph.ShowUndirectedGraphOfNodePair(g)
	graph.ShowUndirectedGraphOfEdge(g)

	graph.ShowDirectedGraphOfNodePair(g)
	graph.ShowDirectedGraphOfEdge(g)
}
```

## License
Copyright (c) 2021 Kazuya yuda.
This software is released under the MIT License, see LICENSE.
https://opensource.org/licenses/mit-license.php
