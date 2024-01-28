package sorting

import (
	"github.com/geovanymds/algorithms/models"
)

func BubbleSort(list *models.LinkedList) {

	var currentNode *models.DoublePointerNode = list.Sentinel.Next

	for i := int64(0); i < list.Length-1; i++ {
		for j := int64(0); j < list.Length-2-i; j++ {
			if currentNode.Id > currentNode.Next.Id {
				currentNode.Next = currentNode.Next.Next
				currentNode.Previous.Next = currentNode.Next.Previous
				currentNode.Next.Previous.Next = currentNode
				currentNode.Next.Previous = currentNode
				currentNode.Previous.Next.Previous = currentNode.Previous
				currentNode.Previous = currentNode.Previous.Next
			} else {
				currentNode = currentNode.Next
			}
		}
		currentNode = list.Sentinel.Next
	}
}
