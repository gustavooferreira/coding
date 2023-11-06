package solutions

import (
	"slices"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkedListFromSlice creates a linked list representing the slice in reverse order.
// A slice []int{1,2,3} will be converted into a linked list like: (3) -> (2) -> (1)
func CreateLinkedListFromSlice(val []int) *ListNode {
	var rootNodePointer *ListNode

	for _, v := range val {
		rootNodePointer = &ListNode{
			Val:  v,
			Next: rootNodePointer,
		}
	}

	return rootNodePointer
}

// IntToSlice returns a slice representing the number passed in.
// Given number 123, it will return a slice like so: []int{1,2,3}
func IntToSlice(num int) (sliceResult []int) {
	for {
		if num == 0 {
			slices.Reverse(sliceResult)
			return sliceResult
		}
		digit := num % 10
		num = num / 10
		sliceResult = append(sliceResult, digit)
	}
}

// GetSliceFromLinkedList returns a slice of ints given a linked list.
func GetSliceFromLinkedList(rootNode *ListNode) []int {
	var result []int

	node := rootNode

	for {
		if node == nil {
			slices.Reverse(result)
			return result
		}

		result = append(result, node.Val)
		node = node.Next
	}
}

// CompareLinkedLists returns true if both linked lists are equal.
func CompareLinkedLists(l1 *ListNode, l2 *ListNode) bool {
	l1Node := l1
	var l1Slice []int

	for {
		if l1Node == nil {
			break
		}
		l1Slice = append(l1Slice, l1Node.Val)
		l1Node = l1Node.Next
	}

	l2Node := l2
	var l2Slice []int

	for {
		if l2Node == nil {
			break
		}
		l2Slice = append(l2Slice, l2Node.Val)
		l2Node = l2Node.Next
	}

	// compare slices
	return slices.Equal(l1Slice, l2Slice)
}
