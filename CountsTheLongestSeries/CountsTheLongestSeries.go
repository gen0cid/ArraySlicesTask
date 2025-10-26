package main

import "fmt"

func main() {

	nums := []int{0, 1, 1, 1, 3, 1, 2, 5, 5, 5, 5, 5}
	CurrentLength := 1
	MaxLength := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			CurrentLength++
		} else {
			if CurrentLength > MaxLength {
				MaxLength = CurrentLength
			}
			CurrentLength = 1

		}
		if CurrentLength > MaxLength {
			MaxLength = CurrentLength
		}
	}
	fmt.Println(MaxLength)
}
