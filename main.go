package main

import (
	"github.com/celso-patiri/DSA/bt"
)

func main() {
	head := &bt.BinaryNode{Value: 1}
	head.Left = &bt.BinaryNode{Value: 2}
	head.Left.Left = &bt.BinaryNode{Value: 3}
	head.Left.Right = &bt.BinaryNode{Value: 4}
	head.Right = &bt.BinaryNode{Value: 5}
	head.Right.Left = &bt.BinaryNode{Value: 6}
	head.Right.Right = &bt.BinaryNode{Value: 7}
}
