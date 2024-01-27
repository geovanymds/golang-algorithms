package models

import "fmt"

type Queue struct {
	Begin *Node
	End   *Node
}

func NewQueue() *Queue {
	return new(Queue)
}

func (queue *Queue) Enqueue(id int32) {
	node := &Node{
		Id:   id,
		Next: nil,
	}

	if queue.Begin == nil {
		queue.Begin = node
		queue.End = node
	} else {
		queue.End.Next = node
		queue.End = node
	}
}

func (queue *Queue) Dequeue() {

	if queue.Begin == nil {
		return
	} else {
		var dequeued *Node = queue.Begin
		queue.Begin = queue.Begin.Next
		dequeued.Next = nil
	}
}

func (queue *Queue) IsEmpty() bool {
	return queue.Begin == nil && queue.End == nil
}

func (queue *Queue) Peek() int32 {
	if queue.Begin != nil {
		return queue.Begin.Id
	} else {
		return -1000
	}
}

func (queue *Queue) PrintBegin() {
	if queue.Begin != nil {
		fmt.Println(queue.Begin.Id)
	}
}
