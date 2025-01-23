package main

import "fmt"


// Graph structure using matrices
type Graph struct {
	vertices int
	matrix [][]int
}

//New Graph initializes a new Graph with the specified  number of vertices

func NewGraph ( vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int , vertices) // Intialize a row with all 0s

	}
	return &Graph {
		vertices : vertices,
		matrix : matrix,
	}

}

// Addedge adds an edge to the graph (directed graph)
// add the feature - weighted graph , and inf for not connected and 0 for no self loop

func (g *Graph) AddEdge(from, to int) {
	if from >= 0 &&  from < g.vertices && to >= 0 && to < g.vertices {
		g.matrix[from][to] = 1  // set the matrices value to 1 to indicate edge 'from' to 'to'
	}
}

func (g *Graph) RemoveEdge(from, to int) {
	if from >= 0 &&  from < g.vertices && to >= 0 && to < g.vertices {
		g.matrix[from][to] = 0 //// set the matrices value to 0 to remove edge 'from' to 'to'
	}
}

// Print the Grapg

func (g *Graph) PrintGraph() {
	fmt.Println (" Print the adjacency matrix : ")
	for i := 0; i < g.vertices; i++ {
		for j:= 0; j < g.vertices; j++ {
			// fmt.Printf("%d", g.matrix[i][j])
			fmt.Print(g.matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func (g *Graph) PrintGraph2() {
	for i := range g.matrix {
		fmt.Println()
		// fmt.Print(g.matrix[i], " ")
		for j := range g.matrix[i] {
			fmt.Print(g.matrix[i][j], " ")
			// fmt.Println()
		}
	}
}


func main() {

	graph := NewGraph(4);

	// Add edges
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	// Print the adjacency matrix representation of the graph
	graph.PrintGraph2()

}