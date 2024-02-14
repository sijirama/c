package main

import (
	"fmt"
    "strings"

)

func closure() {

}

func removeProfanity(str *string) {
    *str= strings.ReplaceAll(*str , "fuck" , "yepa")
}

func square(x int) int { return x * x }
func add(x int) int    { return x + x }

func doMatch(f func(x int) int, list []int) []int {
	var answer []int
	for i := 0; i < len(list); i++ {
		answer = append(answer, f(list[i]))
	}
	return answer
}
