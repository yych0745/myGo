package main

import "fmt"

func main() {
	var x *int
	a := 10
	x = &a
	fmt.Println(x)
	fmt.Println(*x)
	node := &Node{
		val: 10,
		next: &Node{
			val:  20,
			next: nil,
		},
	}
	fmt.Println(node.val)
}

type Node struct {
	val  int
	next *Node
}
