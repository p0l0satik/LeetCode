package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
	Val int
	Next *ListNode
 }
 func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	curr, next := head,head.Next
	for ; next != nil;{
		tmp := next.Next

		next.Next = curr

		curr = next
		next = tmp
	}
	head.Next = nil

	return curr
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
	PrintLL(reverseList(NewLL([]int{1,2,3,4,5})))
	fmt.Println("test")
 }