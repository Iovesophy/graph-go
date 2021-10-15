# graph-go

[![test](https://github.com/Iovesophy/graph-go/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/Iovesophy/graph-go/actions/workflows/test.yml) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/Iovesophy/graph-go/master/LICENSE)

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

	val1 , _ := g.CreateUndirectedVertexesPair()
	val2 , _ := g.CreateUndirectedEdge()
	fmt.Println(val1,val2)
	graph.ShowUndirectedGraph(g)
}
```

https://play.golang.org/p/s0srXwaGh6M

## License
Copyright (c) 2021 Kazuya yuda.
This software is released under the MIT License, see LICENSE.
https://opensource.org/licenses/mit-license.php
