package main

import "fmt"


type Node struct {
	data int
	leftNode *Node
	rightNode *Node
}

func main() {
	var b_tree Node

	b_tree.data = 2
	b_tree.leftNode = nil
	b_tree.rightNode = nil
	fmt.Println(b_tree.data)
	b_tree2 := createNode(6)
	fmt.Println(b_tree2)

}

func createNode(data int) Node {
	var b_tree Node
	b_tree.data = data
	b_tree.leftNode = nil
	b_tree.rightNode = nil
	return b_tree
}
