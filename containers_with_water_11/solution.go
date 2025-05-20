package main

import "fmt"

func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    var max int 
    for st, end := 0, len(height)-1; st < end; {
        curr := (end-st)*min(height[st], height[end])
        if curr > max {
            max = curr
        }
        if height[st] > height[end] {
            end--
        } else {
            st++
        }
    }

    return max
}

func main(){
    fmt.Println(maxArea([]int{1, 8,6,2,5,4,8,3,7}))
    fmt.Println(maxArea([]int{2,2, 2}))
    fmt.Println(maxArea([]int{1, 7,2,5,4,7,3,6}))
    fmt.Println(maxArea([]int{1,1}))


}