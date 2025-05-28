package main

import (
	"fmt"
)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	curr := (l+r) / 2
	if nums[0] == target{
		return 0
	}
	for ;l < r; curr = (l+r) / 2 {
		switch{
		case nums[curr] == target:
			return curr
		case r - l == 1:
			if nums[r] == target{
				return r
			}
			return -1
		case nums[curr] < target:
			l = curr
		case nums[curr] > target:
			r = curr
		}
	}
	return -1
}

func main(){
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(search([]int{-1,0,3,5,9,12}, 12))

	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))



}

