package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRemoveFromSorted(t *testing.T) {
	testCases := []struct {
		origVals []int
		expVals  []int
	}{
		{[]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 3, 3, 3, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = removeDuplicateFromSorted(origList.head)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

func TestRemoveFromUnsorted(t *testing.T) {
	testCases := []struct {
		origVals []int
		expVals  []int
	}{
		{[]int{1, 4, 1, 5, 4, 1, 4, 5, 6, 6}, []int{1, 4, 5, 6}},
		{[]int{5, 4, 3, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{3, 1, 6, 2, 8, 4}, []int{3, 1, 6, 2, 8, 4}},
		{[]int{5, 3, 5, 3, 4, 4}, []int{5, 3, 4}},
		{[]int{5, 5, 5, 5, 5, 3, 4, 4, 4, 3, 4, 4, 4}, []int{5, 3, 4}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = removeDuplicateFromUnsorted(origList.head)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

func TestSwapTwoGivenElements(t *testing.T) {
	testCases := []struct {
		origVals []int
		x        int
		y        int
		expVals  []int
	}{
		{[]int{5, 10, 2, 6, 11, 9}, 5, 5, []int{5, 10, 2, 6, 11, 9}},
		{[]int{5, 10, 2, 6, 11, 9}, 5, 6, []int{6, 10, 2, 5, 11, 9}},
		{[]int{5, 10, 2, 6, 11, 9}, 5, 9, []int{9, 10, 2, 6, 11, 5}},
		{[]int{5, 10, 2, 6, 11, 9}, 2, 6, []int{5, 10, 6, 2, 11, 9}},
		{[]int{5, 10, 2, 6, 11, 9}, 6, 2, []int{5, 10, 6, 2, 11, 9}},
		{[]int{5, 10, 2, 6, 11, 9}, 9, 10, []int{5, 9, 2, 6, 11, 10}},
		{[]int{5, 10, 2, 6, 11, 9}, 2, 8, []int{5, 10, 2, 6, 11, 9}},
		{[]int{1}, 1, 2, []int{1}},
		{[]int{}, 1, 2, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = swapTwoGivenElements(origList.head, tc.x, tc.y)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

func TestPairWiseSwapElements(t *testing.T) {
	testCases := []struct {
		origVals []int
		expVals  []int
	}{
		{[]int{5, 10, 2, 6, 11, 9}, []int{10, 5, 6, 2, 9, 11}},
		{[]int{5, 10, 2, 6, 11}, []int{10, 5, 6, 2, 11}},
		{[]int{5, 10}, []int{10, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = pairWiseSwap(origList.head)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

func TestMoveLastElementToFront(t *testing.T) {
	testCases := []struct {
		origVals []int
		expVals  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3}, []int{3, 1, 2}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = moveLastElementToFront(origList.head)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

func TestIntersectionOfTwoLists(t *testing.T) {
	testCases := []struct {
		origVals1 []int
		origVals2 []int
		expVals   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{}},
		{[]int{1, 3, 5}, []int{1, 3, 5}, []int{1, 3, 5}},
		{[]int{1, 3, 5}, []int{}, []int{}},
		{[]int{}, []int{1, 2, 3}, []int{}},
	}

	for _, tc := range testCases {
		origList1, origList2, expList := createList(tc.origVals1), createList(tc.origVals2), createList(tc.expVals)
		thirdList := &LinkedList{head: intersectionOfTwoSortedLists(origList1.head, origList2.head)}
		if !thirdList.equals(expList) {
			t.Fatal("Two lists differ\n", thirdList, "\n", expList)
		}
	}
}

// Problem no. 24 - https://www.geeksforgeeks.org/segregate-even-and-odd-elements-in-a-linked-list/
func TestSegregateEvenAndOdd(t *testing.T) {
	testCases := []struct {
		origVals []int
		expVals  []int
	}{
		{[]int{17, 15, 8, 12, 10, 5, 4, 1, 7, 6}, []int{8, 12, 10, 4, 6, 17, 15, 5, 1, 7}},
		{[]int{8, 12, 10}, []int{8, 12, 10}},
		{[]int{1, 3, 5}, []int{1, 3, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		origList, expList := createList(tc.origVals), createList(tc.expVals)
		origList.head = segregateEvenAndOdd(origList.head)
		if !origList.equals(expList) {
			t.Fatal("Two lists differ\n", origList, "\n", expList)
		}
	}
}

// Problem no. 25 - https://www.geeksforgeeks.org/print-reverse-of-a-linked-list-without-actually-reversing/
func TestPrintInReverse(t *testing.T) {
	testCases := []struct {
		origVals []int
		exp      []byte
	}{
		{[]int{1, 2, 3, 4}, []byte{4, 3, 2, 1}},
	}

	buf := &bytes.Buffer{}
	for _, tc := range testCases {
		origList := createList(tc.origVals)
		printInReverse(origList.head, buf)
		if reflect.DeepEqual(buf.Bytes(), tc.exp) {
			t.Fatal("Two lists differ\n", buf.String(), "\n", tc.exp)
		}
	}
}

// helpers
func createList(values []int) *LinkedList {
	list := &LinkedList{}
	for _, val := range values {
		list.addValue(val)
	}
	return list
}
