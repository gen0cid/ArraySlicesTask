package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Анна"
	s = strings.ToLower(s)
	r := []rune(s)
	p := true

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			p = false
			break
		}

	}

	if p == false {
		fmt.Println("Не палиндром")
	} else {
		fmt.Println("Палиндром")

	}
}
