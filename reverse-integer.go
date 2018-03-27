package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	k := 1
	if x < 0 {
		x *= -1
		k = -1
	}

	res := 0
	for x > 0 {
		i := x % 10
		res = res*10 + i
		x = (x - i) / 10
		if res > math.MaxInt32 {
			return 0
		}
	}

	return k * res
}

func main() {
	n1 := 12345
	n2 := 2147483648
	n3 := 9876543219
	fmt.Println(n1, "=>", reverse(n1))
	fmt.Println(n2, "=>", reverse(n2))
	fmt.Println(n3, "=>", reverse(n3))
}
