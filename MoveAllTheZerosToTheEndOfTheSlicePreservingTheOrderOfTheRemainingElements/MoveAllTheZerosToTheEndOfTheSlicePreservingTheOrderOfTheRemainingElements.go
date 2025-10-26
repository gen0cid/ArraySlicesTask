package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	nums1 := []int{}
	zeros := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros++
		} else {
			nums1 = append(nums1, nums[i])
		}

	}
	for j := 0; j < zeros; j++ {
		nums1 = append(nums1, 0)

	}
	fmt.Println(nums1)
}
