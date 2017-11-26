package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)

	found := false
	for i, v := range nums {
		if v == num {
			fmt.Printf("%d found at index %d\n", num, i)
			found = true
		}
	}
	if !found {
		fmt.Printf("%d not found\n", num)
	}
}
func main() {
	find(1, 2, 3, 4, 5, 6)
	find(2, 3, 4, 5, 6, 2, 7)
}
