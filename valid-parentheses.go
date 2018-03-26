package main

import (
	"fmt"
	"strings"
)

var startings = "([{"
var endings = ")]}"

func endingToStarting(s string) string {
	i := strings.Index(endings, s)
	return startings[i : i+1]
}

func isValid(s string) bool {
	acc := ""

	for _, char := range s {
		sub := string(char)
		if strings.Contains(startings, sub) {
			acc += sub
		} else if strings.Contains(endings, sub) {
			if strings.HasSuffix(acc, endingToStarting(sub)) {
				acc = acc[0 : len(acc)-1]
			} else {
				return false
			}
		}
	}

	return len(acc) == 0
}

func main() {
	s1 := "()[]{}"
	s2 := "([)]"
	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
}
