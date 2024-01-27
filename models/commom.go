package models

type Node struct {
	Id   int32
	Next *Node
}

type DoublePointerNode struct {
	Id       int32
	Next     *DoublePointerNode
	Previous *DoublePointerNode
}
