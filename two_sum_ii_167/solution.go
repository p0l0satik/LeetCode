package main

import "fmt"
func twoSum(numbers []int, target int) []int {
	
	for st, end := 0, len(numbers)-1; st < end; {
		switch sum := numbers[st] + numbers[end]; {
		case sum > target:
			end--;
		case sum < target:
			st++;
		default:
			return []int{st +1, end + 1}
		}
	}
	return []int{}
}

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 11, 15}, 9))
	fmt.Println(twoSum([]int{1, 3, 4, 5,7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}