package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := [5]int{2, 2, 1, 1, 2}
	c := math.Ceil(float64(len(nums)) / float64(2))
	sort.Ints(nums[:])
	fmt.Println(nums)
	var a float64 = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			a++
		} else if nums[i] != nums[i+1] {
			a = 1
		}
		if (len(nums)%2 == 0) && a > c {
			fmt.Println(nums[i])
			break
		} else if (len(nums)%2 != 0) && a == c {
			fmt.Println(nums[i])
			break
		}

	}
}
