package main

import (
	"fmt"
	"reflect"
)

func main() {

	switch finger := 10; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinkie")
	default:
		fmt.Println("Unknown finger")
	}

	for index, runeValue := range "Hello World!" {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	fmt.Println(reflect.TypeOf("Hello World!!"))

	switch letter := "d"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

}
