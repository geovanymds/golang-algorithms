package main

import (
	"github.com/geovanymds/algorithms/models"
)

func main() {
	var list *models.LinkedList = models.NewList()
	list.Insert(95)
	list.Insert(9)
	list.Insert(11)
	list.Insert(87)
	list.Insert(11)
	list.Insert(95)
	list.Insert(21)

	list.RemoveAll(95)

	list.Display()
}
