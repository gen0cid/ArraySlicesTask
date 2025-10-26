package main

import "fmt"

func main() {
	nums := []int{1, 2, 34, 4, 6, 7, 8, 9, 54, 6, 8, 9, 5, 4}
	nums1 := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			nums1 = append(nums1, nums[i])

		}
	}
	fmt.Println(nums1)
}
