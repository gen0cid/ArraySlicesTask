package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	c := "А, роза упала на лапу Азора!"
	c = strings.ToLower(c)
	ClassicRune := []rune(c)
	var clean []rune
	f := true

	for _, r := range ClassicRune {
		if unicode.IsLetter(r) {
			clean = append(clean, r)
		}
	}
	seredina := len(clean) / 2

	for i := 0; i < seredina; i++ {
		if clean[i] != clean[len(clean)-1-i] {
			f = false
			break
		}
	}
	if f == false {
		fmt.Println("не палиндром")
	} else {
		fmt.Println("Палиндром")

	}
}
