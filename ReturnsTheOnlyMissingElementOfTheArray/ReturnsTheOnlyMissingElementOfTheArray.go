package main

import "fmt"

func main() {
	nums := []int{1, 3, 6, 2, 5, 4, 8, 7}
	var sum float64 = (float64(len(nums)) + 1) * (float64(len(nums))) / 2
	var c float64
	for i := 0; i < len(nums); i++ {
		c = c + float64(nums[i])
	}

	fmt.Println(sum - c)

}
