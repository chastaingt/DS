package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", key)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: key})
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}

	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func contains(list []*Vertex, key int) bool {
	for _, v := range list {
		if key == v.key {
			return true
		}
	}

	return false
}

func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(6, 2)
	graph.AddEdge(3, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)

	graph.Print()
}
