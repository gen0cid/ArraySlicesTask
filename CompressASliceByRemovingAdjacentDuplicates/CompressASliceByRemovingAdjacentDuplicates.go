package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 2, 3, 4, 4, 5, 5, 5, 6}
	nums1 := []int{}
	nums1 = append(nums1, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums1 = append(nums1, nums[i])
		}
	}
	fmt.Println(nums1)
}
