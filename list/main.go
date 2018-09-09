package main

import (
	"fmt"
)

func main() {
	list := &LinkedList{}

	list.addValue(1)
	list.addValue(2)
	list.addValue(3)
	list.addValue(99)
	list2 := &LinkedList{}
	list2.addValue(10)
	list2.addValue(11)
	list2.tail.next = list.tail
	list.addValue(4)
	list.addValue(5)

	fmt.Println("Intersection point is ", intersectionPointOfTwoLists(list.head, list2.head).val)

}
