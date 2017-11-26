package main

import _ "fmt"
import "fmt"

func main() {
	for i, j := 0, 10; i <= 10 && j <= 19; i, j = i+1, j+1 {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d * %d = %d", i, j, i*j)
		fmt.Println()
	}
}
