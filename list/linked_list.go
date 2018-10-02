package main

import (
	"fmt"
	"io"
	"strings"
)

// Node denotes an item in the list
type Node struct {
	val  int
	next *Node
}

// LinkedList denotes a linked list
type LinkedList struct {
	head,
	tail *Node
}

func (list *LinkedList) String() string {
	result := []string{}
	for curr := list.head; curr != nil; curr = curr.next {
		result = append(result, fmt.Sprintf("%d", curr.val))
	}
	return strings.Join(result, "->")
}

func (list *LinkedList) equals(otherList *LinkedList) bool {
	curr1, curr2 := list.head, otherList.head
	for ; curr1 != nil && curr2 != nil && curr1.val == curr2.val; curr1, curr2 = curr1.next, curr2.next {
	}
	return curr1 == nil && curr2 == nil
}

func (list *LinkedList) addValue(val int) {
	node := &Node{val: val}
	if list.head == nil {
		list.head = node
		list.tail = list.head
		return
	}

	list.tail.next = node
	list.tail = node
}

func (list *LinkedList) addValueAt(index, val int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	if index < 0 {
		fmt.Println("Non-negative index reqd")
		return
	}

	if index == 0 {
		node := &Node{val: val, next: list.head}
		list.head = node
		return
	}

	currIndex := 1
	for curr := list.head; curr != nil; curr = curr.next {
		if currIndex == index {
			node := &Node{val: val, next: curr.next}
			curr.next = node
			return
		}
		currIndex++
	}

	fmt.Println("List smaller than index")
}

func reverseIterative(node *Node) *Node {
	if node == nil || node.next == nil {
		fmt.Println("List contains one or zero nodes")
		return node
	}

	prev := node
	curr := node.next
	prev.next = nil
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func reverseRecursive(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}

	p := reverseRecursive(node.next)
	node.next.next = node
	node.next = nil
	return p
}

func findLength(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + findLength(node.next)
}

func findLengthIter(node *Node) int {
	i := 0
	for ; node != nil; i, node = i+1, node.next {
	}
	return i
}

func isPresent(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	}

	return isPresent(node.next, val)
}

func isPresentIter(node *Node, val int) bool {
	for curr := node; curr != nil; curr = curr.next {
		if curr.val == val {
			return true
		}
	}

	return false
}

func nThNode(node *Node, n int) *Node {
	if n < 0 || node == nil {
		return nil
	}

	if n == 0 {
		return node
	}

	return nThNode(node.next, n-1)
}

func nThNodeFromEndIter(node *Node, n int) *Node {
	if n < 0 || node == nil {
		return nil
	}

	curr := node
	for i := 0; node != nil; i, node = i+1, node.next {
		if i < n {
			continue
		}
		curr = curr.next
	}

	return curr
}

func middleOfList(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}

	fast, slow := node, node
	for ; fast != nil && fast.next != nil; slow, fast = slow.next, fast.next.next {
	}

	return slow
}

func frequencyOfVal(node *Node, val int) (count int) {
	for ; node != nil; node = node.next {
		if node.val == val {
			count++
		}
	}
	return
}

func detectLoop(node *Node) bool {
	if node == nil || node.next == nil {
		return false
	}

	slow, fast := node, node.next
	for ; fast != nil && fast.next != nil && fast != slow; slow, fast = slow.next, fast.next.next {
	}

	if fast == slow {
		return true
	}
	return false
}

func detectLengthOfLoop(node *Node) int {
	if node == nil || node.next == nil {
		return 0
	}

	slow, fast := node, node.next
	for ; fast != nil && fast.next != nil && fast != slow; slow, fast = slow.next, fast.next.next {
	}

	if fast != slow {
		return 0
	}

	length := 1
	for fast = slow.next; fast != slow; fast = fast.next {
		length++
	}

	return length
}

func isPalindrome(node *Node) bool {
	if node == nil {
		return false
	}
	if node.next == nil {
		return true
	}

	slow, fast := node, node.next
	for ; fast != nil && fast.next != nil; slow, fast = slow.next, fast.next.next {

	}

	mid1, mid2 := slow, slow

	if fast != nil {
		mid2 = mid2.next
		if mid1.val != mid2.val {
			return false
		}
	}

	head2, isPal := reverseList(mid2), true

	for curr1, curr2 := node, head2; curr1 != mid1 && curr2 != mid2; curr1, curr2 = curr1.next, curr2.next {
		if curr1.val != curr2.val {
			isPal = false
			break
		}
	}

	mid2 = reverseList(head2)

	return isPal
}

func reverseList(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}

	p := reverseList(node.next)
	node.next.next = node
	node.next = nil
	return p
}

func removeDuplicateFromSorted(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}
	p := node

	for curr1, curr2 := node, node.next; curr2 != nil; curr2 = curr2.next {
		if curr1.val == curr2.val {
			curr1.next = curr2.next
		} else {
			curr1 = curr2
		}

	}

	return p
}

func removeDuplicateFromUnsorted(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}
	visitedVals := map[int]struct{}{node.val: struct{}{}}

	p, curr1, curr2 := node, node, node.next
	for ; curr2 != nil; curr2 = curr2.next {
		if _, ok := visitedVals[curr2.val]; ok {
			curr1.next = curr2.next

		} else {
			visitedVals[curr2.val] = struct{}{}
			curr1 = curr1.next
		}
	}

	return p
}

func swapTwoGivenElements(node *Node, x, y int) *Node {
	if node == nil || node.next == nil || x == y {
		return node
	}

	var xPrev, yPrev *Node
	xCurr := node
	for ; xCurr != nil; xCurr = xCurr.next {
		if xCurr.val == x {
			break
		}
		xPrev = xCurr
	}

	yCurr := node
	for ; yCurr != nil; yCurr = yCurr.next {
		if yCurr.val == y {
			break
		}
		yPrev = yCurr
	}

	// if head.val == x
	if xPrev == nil {
		head := yPrev.next
		xNext := node.next
		node.next = head.next
		yPrev.next = node
		head.next = xNext
		return head
	}

	// if head.val == y
	if yPrev == nil {
		head := xPrev.next
		xPrev.next = node
		yNext := node.next
		node.next = head.next
		head.next = yNext
		return head
	}

	if xPrev.next == nil || yPrev.next == nil {
		fmt.Println("x or y not found")
		return node
	}

	if xPrev.next == yPrev {
		xPrev.next = yPrev.next
		temp := yPrev.next.next
		yPrev.next.next = yPrev
		yPrev.next = temp
		return node
	}

	if yPrev.next == xPrev {
		yPrev.next = xPrev.next
		temp := xPrev.next.next
		xPrev.next.next = xPrev
		xPrev.next = temp
		return node
	}

	xNode, yNode := xPrev.next, yPrev.next
	xPrev.next = yNode
	yNext := yNode.next
	yNode.next = xNode.next
	yPrev.next = xNode
	xNode.next = yNext
	return node
}

func pairWiseSwap(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}

	head := node.next
	temp := head.next
	head.next = node
	node.next = pairWiseSwap(temp)
	return head
}

func moveLastElementToFront(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}

	head, secondLast := node, node
	for ; secondLast.next.next != nil; secondLast = secondLast.next {
	}
	secondLast.next.next = head
	head = secondLast.next
	secondLast.next = nil
	return head
}

func intersectionOfTwoSortedLists(node1, node2 *Node) *Node {
	var head3, curr3 *Node
	if node1 == nil || node2 == nil {
		return head3
	}

	for curr1, curr2 := node1, node2; curr1 != nil && curr2 != nil; {
		if curr1.val == curr2.val {
			temp := &Node{val: curr1.val}
			if head3 == nil {
				head3 = temp
				curr3 = head3
			} else {
				curr3.next = temp
				curr3 = curr3.next
			}
			curr1, curr2 = curr1.next, curr2.next
		} else if curr1.val < curr2.val {
			curr1 = curr1.next
		} else {
			curr2 = curr2.next
		}

	}

	return head3
}

func intersectionPointOfTwoLists(node1, node2 *Node) *Node {
	if node1 == nil || node2 == nil {
		return nil
	}
	last1, lengnth1 := node1, 1
	for ; last1.next != nil; last1 = last1.next {
		lengnth1++
	}
	last1.next = node1
	defer func() {
		last1.next = nil
	}()

	curr1, curr2 := node2, node2
	for ; lengnth1 > 0; curr1, lengnth1 = curr1.next, lengnth1-1 {
	}

	for ; curr1 != curr2; curr1, curr2 = curr1.next, curr2.next {
	}
	return curr1
}

// Problem no. 24 - https://www.geeksforgeeks.org/segregate-even-and-odd-elements-in-a-linked-list/
func segregateEvenAndOdd(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	var evenHead, evenCurr, oddHead, oddCurr, next *Node
	for curr := head; curr != nil; curr = next {
		next = curr.next
		if curr.val%2 == 0 {
			if evenHead == nil {
				evenHead = curr
			} else {
				evenCurr.next = curr
			}
			evenCurr = curr
		} else {
			if oddHead == nil {
				oddHead = curr
			} else {
				oddCurr.next = curr
			}
			oddCurr = curr
		}
	}

	if evenCurr == nil {
		evenHead = oddHead
	} else {
		evenCurr.next = oddHead
	}
	if oddCurr != nil {
		oddCurr.next = nil
	}
	return evenHead
}

// Problem no. 26 - https://www.geeksforgeeks.org/print-reverse-of-a-linked-list-without-actually-reversing/
func printInReverse(head *Node, w io.Writer) {
	if head == nil {
		return
	}
	if head.next == nil {
		fmt.Fprint(w, head)
	}
	printInReverse(head.next, w)
}

// // Problem no. 28 - https://www.geeksforgeeks.org/merge-two-sorted-linked-lists-such-that-merged-list-is-in-reverse-order/
// func mergeListInReverse(head1, head2 *Node) *Node {
// 	var last, next1, next2 *Node

// 	for curr1, curr2 := head1, head2; curr1 != nil || curr2 != nil; {
// 		if curr1 != nil&curr2 != nil {
// 			if curr1.val < curr2.val {
// 				next1 = curr1.next

// 			}

// 		}
// 	}
// }
