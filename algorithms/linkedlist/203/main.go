package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	current := head
	var previous ListNode

	for current != nil {
		// print
		fmt.Println(current.Val)
		if current.Val == val {
			temp := current
			current = current.Next
			fmt.Println("prev is", previous.Val)
			//fmt.Println("curr is", current.Val)
			previous.Next = current
			//fmt.Println("new prev next val", previous.Next.Val)
			//current.Next=current.Next.Next
			temp.Next = nil
			if head == temp {
				head = current
			}
			continue
		}
		if current.Next == nil {
			break
		}
		previous = *current
		fmt.Println("prev is outside", previous.Val)
		current = current.Next
	}
	//if current == nil { return current }
	return head

}
