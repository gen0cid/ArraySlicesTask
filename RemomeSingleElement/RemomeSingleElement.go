package main

import "fmt"

/*func main() {
	c := []int{2, 4, 2, 4, 5, 6, 7, 6, 7}
	for i := 0; i < len(c); i++ {
		for j := 0; j < i-1; j++ {
			if j == i {
				for k := i + 1; k < len(c); k++ {
					if c[k] == c[i] {
						break
					} else {
						fmt.Println()
					}

				}
			}
			if c[j] == c[i] {

				break
			}
		}

	}
}
*/

func main() {
	c := []int{3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3}
	i := 0
Restart:
	for j := i + 1; j < len(c); j++ {

		if c[i] == c[j] {
			part2 := c[i+1 : j]
			part3 := c[j+1:]

			c = append(part2, part3...)
			goto Restart
		}

	}

	fmt.Println(c[0])
}
