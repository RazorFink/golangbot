package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Print("\n")
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Print("\n")
}

func main() {
	s := "Hello World!"
	printBytes(s)
	printChars(s)
	s = "Señor"
	printBytes(s)
	printChars(s)
}
