package main

import (
	"fmt"
	"unicode/utf8"
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

func printRunes(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Print("\n")
}

func printRunesAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func printByteSlice() {
	bytesSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(bytesSlice)
	fmt.Println(str)
}

func printRuneSlice() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

func printAllRunes() {
	for r := 0x1f600; r <= 0x1ffff; r++ {
		fmt.Println(string(r))
	}
}

func printLen(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	s := "Hello World!"
	printBytes(s)
	printChars(s)
	printRunes(s)
	printRunesAndBytes(s)
	s = "Señor"
	printBytes(s)
	printChars(s)
	printRunes(s)
	printRunesAndBytes(s)
	printByteSlice()
	printRuneSlice()
	// printAllRunes()
	word1 := "Pets"
	word2 := "Señor"
	fmt.Printf("bytes in %s: %d\n", word1, len(word1))
	fmt.Printf("bytes in %s: %d\n", word2, len(word2))
	printLen(word1)
	printLen(word2)
	fmt.Println(mutate([]rune(word1)))
	fmt.Println(word1)
	fmt.Println(mutate([]rune(word2)))
	fmt.Println(word2)
}
