package main

import "fmt"

func main() {
	c := 1
	a := []int{1, 1, 2, 2, 3, 4, 4}
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			c++
		} else if i == len(a)-2 {
			fmt.Println(a[len(a)-1])
			break
		} else if (a[i] != a[i+1]) && (c != 1) {
			c = 1
		} else if (a[i] != a[i+1]) && (c == 1) {
			fmt.Println(a[i])
			break
		}

	}

}
