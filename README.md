# graph-go

## sample code

```Go
package main

import (
	"fmt"
	graph "github.com/Iovesophy/graph-go"
)

func main() {
	samplePlace := []graph.Node{
		graph.Node{ID: 1, Element: "place-A"},
		graph.Node{ID: 2, Element: "place-B"},
		graph.Node{ID: 3, Element: "place-C"},
		graph.Node{ID: 4, Element: "place-D"},
	}
	sampleStation := []graph.Node{
		graph.Node{ID: 5, Element: "station-A"},
		graph.Node{ID: 6, Element: "station-B"},
	}
	g := &graph.Glink{
		NodeData: samplePlace,
		BaseData: sampleStation,
	}

	fmt.Println(g.CreateVertexesPair())
	fmt.Println(g.CreateEdge())

	graph.DoGraph(g)
}
```

## License
Copyright (c) 2021 Kazuya yuda.
This software is released under the MIT License, see LICENSE.
https://opensource.org/licenses/mit-license.php

## Authors
kazuya yuda.
