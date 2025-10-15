package main

import "fmt"

func main() {
	nums := [4][4]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
	}
	rows := len(nums)
	columns := len(nums[0])
	var c int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if nums[i][j] == 1 {
				if j == 0 || (nums[i][j-1] == 0) {
					c++
				}
				if j == columns-1 || (nums[i][j+1] == 0) {
					c++
				}
				if i == 0 || (nums[i-1][j] == 0) {
					c++
				}
				if i == rows-1 || (nums[i+1][j] == 0) {
					c++
				}
			}

		}

	}
	fmt.Println(c)
}
