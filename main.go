package main

import "github.com/eddie/gograph"
import "fmt"

func main() {

	g := graph.CreateGraph(false)

	g.ReadGraph("data/graph.txt")

	fmt.Println("Shortest Path between vertex 6 and 4\n")

	pathFunc := func(x int) {
		fmt.Printf("%d -> ", x)
	}
	g.FindPath(6, 4, pathFunc)

	fmt.Println("")
}
