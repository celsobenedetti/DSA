package main

import (
	"fmt"

	"github.com/celso-patiri/DSA/algorithms"
	"github.com/celso-patiri/DSA/structures"
)

func main() {
	head := &structures.BinaryNode{Value: 1}
	head.Left = &structures.BinaryNode{Value: 2}
	head.Left.Left = &structures.BinaryNode{Value: 3}
	head.Left.Right = &structures.BinaryNode{Value: 4}
	head.Right = &structures.BinaryNode{Value: 5}
	head.Right.Left = &structures.BinaryNode{Value: 6}
	head.Right.Right = &structures.BinaryNode{Value: 7}

	fmt.Println(algorithms.PostOrderSearch(head))
	fmt.Println(algorithms.BFS(head, 10))
}
