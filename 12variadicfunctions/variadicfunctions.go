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

func change(s ...string) {
	s = append(s, "playground")
	s[0] = "Go"
	fmt.Println(s)
}

func main() {
	find(1, 2, 3, 4, 5, 6)
	find(2, 3, 4, 5, 6, 2, 7)
	s := []int{8, 7, 6, 5, 4, 3, 7}
	find(7, s...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
