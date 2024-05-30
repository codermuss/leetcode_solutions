package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//** first list
	var ln3 ListNode = ListNode{Val: 4, Next: nil}
	var ln2 ListNode = ListNode{Val: 2, Next: &ln3}
	var ln1 ListNode = ListNode{Val: 1, Next: &ln2}

	//** second list
	var ln6 ListNode = ListNode{Val: 4, Next: nil}
	var ln5 ListNode = ListNode{Val: 3, Next: &ln6}
	var ln4 ListNode = ListNode{Val: 1, Next: &ln5}

	var result *ListNode = mergeTwoLists(&ln1, &ln4)
	for result.Next != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}
