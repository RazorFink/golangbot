package main

import (
	"fmt"
)

func declaringPointers() {
	b := 255
	var a = &b // a is *int, inferred by &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("Address of b is", a)
	fmt.Println("Address of a is", &a)
}
func zeroValue() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

func dereferencing() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", *a)
}

func change(val *int) {
	*val = 55
}
func passing() {

	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a before function call is", a)
}

// Not idiomatic, don't do this
func modifyPointer(arr *[3]int) {
	// (*arr)[0] = 90
	arr[0] = 90
}

// Not idiomatic, don't do this
func arrayPointerArg() {
	a := [3]int{89, 90, 91}
	modifyPointer(&a)
	fmt.Println(a)
}

// Use slices instead of array pointers
func modifySlice(sls []int) {
	sls[0] = 90
}

func arraySliceArg() {
	a := [3]int{89, 90, 91}
	modifySlice(a[:])
	fmt.Println(a)
}

func main() {
	declaringPointers()
	zeroValue()
	dereferencing()
	passing()
	arrayPointerArg()
	arraySliceArg()
}
