package models

import "fmt"

type Stack struct {
	Top *Node
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) GetTop() *Node {
	return stack.Top
}

func (stack *Stack) PrintTop() {
	if stack.Top != nil {
		fmt.Println(stack.Top.Id)
	}
}

func (stack *Stack) Push(id int32) {
	var newNode = Node{
		Id:   id,
		Next: stack.GetTop(),
	}

	stack.Top = &newNode
}

func (stack *Stack) Pop() {
	if stack.Top.Next != nil {
		var poppedNode *Node = stack.Top
		stack.Top = stack.Top.Next
		fmt.Println("Popped node id: ", poppedNode.Id)
		poppedNode = nil
	} else {
		stack.Top = nil
	}
}
