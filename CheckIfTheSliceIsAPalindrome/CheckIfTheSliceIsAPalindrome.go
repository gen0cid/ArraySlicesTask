package main

import (
	"fmt"
)

func main() {
	a := true
	nums := []int{1, 2, 3, 2, 1}
	seredina := (len(nums)) / 2
	for i := 0; i < seredina; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			a = false
			break
		}
	}
	if a == true {
		fmt.Println("палиндром")
	} else {
		fmt.Println("не палиндром")

	}
}
