package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	p1, p2 := head, head
    for ; p1.Next != nil; p1 = p1.Next {
		if n == 0 {
			p2 = p2.Next
		} else {
			n -= 1
		}
	} 
	if p2 == head && n > 0 {
		return p2.Next
	}
	p2.Next = p2.Next.Next

	return head
}

func NewLL(arr []int) *ListNode{
	head := &ListNode{}
	curr := head
	for i, a := range arr {
		curr.Val = a
		if len(arr) - 1 > i {
			curr.Next = &ListNode{}
		}
		curr = curr.Next
	}
	return head
}

func PrintLL(head *ListNode) *ListNode{
	for ; head != nil; head = head.Next{
		fmt.Printf("%d ", head.Val)
	}
	fmt.Println()
	return head
}

func main(){
    PrintLL(removeNthFromEnd(NewLL([]int{1, 2}), 2))
}