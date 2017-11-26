package main

import _ "fmt"
import "fmt"

func main() {
	i := 0
	for ; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
		i += 1
	}
}
