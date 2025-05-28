package main

import "fmt"

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

func Append(root *TreeNode, v int){
	for ;; {
		if v < root.Val {
			if root.Left == nil {
				root.Left = &TreeNode{
					Val:  v,
				}
				return
			}
			root = root.Left
		} else {
			if root.Right == nil {
				root.Right = &TreeNode{
					Val:  v,
				}
				return
			}
			root = root.Right
		}
	}
}

func PrintTree(root *TreeNode){
	toProcess := []*TreeNode{root}
	ptr := 0
	for ; ptr < len(toProcess); ptr++ {
		node := toProcess[ptr]
		if node.Left != nil {
			toProcess = append(toProcess, node.Left)
		}
		if node.Right != nil {
			toProcess = append(toProcess, node.Right)
		}
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

 
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	toProcess := []*TreeNode{root}
	ptr := 0
	for ; ptr < len(toProcess); ptr++ {
		node := toProcess[ptr]
		if node.Left != nil {
			toProcess = append(toProcess, node.Left)
		}
		if node.Right != nil {
			toProcess = append(toProcess, node.Right)
		}
		temp :=  node.Right
		node.Right = node.Left
		node.Left = temp
	}
	return root
}

func main(){
	t1 := &TreeNode{ Val: 4}
	for _, v := range []int{2,7,1,3,6,9}{
		Append(t1, v)
	}
	PrintTree(t1)
	invertTree(t1)
	PrintTree(t1)
}