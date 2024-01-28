package models

import "fmt"

type LinkedList struct {
	Sentinel *DoublePointerNode
	Length   int64
}

func NewList() *LinkedList {
	var sentinel *DoublePointerNode = &DoublePointerNode{Id: -999999}
	sentinel.Next = sentinel
	sentinel.Previous = sentinel
	return &LinkedList{sentinel, 0}
}

func (list *LinkedList) Insert(id int32) {
	var node *DoublePointerNode = &DoublePointerNode{Id: id}

	node.Next = list.Sentinel.Next
	list.Sentinel.Next.Previous = node
	list.Sentinel.Next = node
	node.Previous = list.Sentinel
	list.Length++
}

func (list *LinkedList) Display() {
	node := list.Sentinel.Next
	for node.Id != -999999 {
		fmt.Println(node.Id)
		node = node.Next
	}
}

func (list *LinkedList) Search(id int32) *DoublePointerNode {
	var node *DoublePointerNode = list.Sentinel.Next

	for node.Id != -999999 && node.Id != id {
		node = node.Next
	}

	return node
}

func (list *LinkedList) RemoveAll(id int32) int64 {
	count := int64(0)
	var node *DoublePointerNode = list.Search(id)

	for node.Id != -999999 {
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
		node.Next = nil
		node.Previous = nil
		node = list.Search(id)
		count++
	}

	list.Length -= count

	return count
}
