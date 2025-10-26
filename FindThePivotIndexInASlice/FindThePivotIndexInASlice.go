package main

import "fmt"

func main() {
	padenie := 0
	nums := []int{4, 5, 6, 7}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			padenie = i + 1
			break
		}
	}
	fmt.Println(padenie)
}
