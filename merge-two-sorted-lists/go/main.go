package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	listIterator := list

	for listIterator != nil {
		fmt.Printf("%v -> ", listIterator.Val)
		listIterator = listIterator.Next
	}
	fmt.Printf("nil")

	fmt.Println()
}

// mergeTwoLists merges two sorted linked lists into one sorted linked list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	outputList := &ListNode{}
	outputListHead := outputList
	l1Ptr := list1
	l2Ptr := list2
	for {
		if l1Ptr == nil {
			outputList.Val = l2Ptr.Val
			l2Ptr = l2Ptr.Next
		} else if l2Ptr == nil {
			outputList.Val = l1Ptr.Val
			l1Ptr = l1Ptr.Next
		}

		if l1Ptr != nil && l2Ptr != nil && l1Ptr.Val > l2Ptr.Val {
			outputList.Val = l2Ptr.Val
			l2Ptr = l2Ptr.Next
		} else if l1Ptr != nil && l2Ptr != nil && l1Ptr.Val <= l2Ptr.Val {
			outputList.Val = l1Ptr.Val
			l1Ptr = l1Ptr.Next
		}

		if l1Ptr == nil && l2Ptr == nil {
			break
		}

		outputList.Next = &ListNode{}
		outputList = outputList.Next
	}

	return outputListHead
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	// list1 := &ListNode{
	// 	Val: -9,
	// 	Next: &ListNode{
	// 		Val:  3,
	// 		Next: nil,
	// 	},
	// }

	// list2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val:  7,
	// 		Next: nil,
	// 	},
	// }

	// list1 := &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	// list2 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }

	// list1 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }

	// list2 := &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 5,
	// 		Next: &ListNode{
	// 			Val:  7,
	// 			Next: nil,
	// 		},
	// 	},
	// }

	mergedList := mergeTwoLists(list1, list2)
	printList(list1)
	printList(list2)
	printList(mergedList)
}
