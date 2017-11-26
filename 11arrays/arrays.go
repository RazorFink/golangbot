package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a [3]int
	a[0] = 12
	a[1] = 25
	a[2] = 32
	fmt.Println(a)

	b := [5]int{3, 7, 23, 31, 1}
	fmt.Println(b)

	c := [...]int{2, 4, 6, 8}
	fmt.Println(c)

	d := [...]string{"miata", "3", "5", "cosmo", "mpv"}
	e := d
	e[0] = "mx5"

	fmt.Println(d)
	fmt.Println(e)

	// array size is part of the type
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))

	for i := 0; i < 26; i++ {
		fmt.Printf("%c\n", 'a'+i)
	}

	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))
	fmt.Println(len(d))
	fmt.Println(len(e))

	fmt.Println(len("Hello World!"))

	f := "Hello World!"[3:7]
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	g := e[1:3]
	fmt.Println(g)
	fmt.Println(reflect.TypeOf(g))

	h := [...]string{"Ferrari", "Honda", "Ford"}
	i := [...]string{"Toyota", "Fiat"}
	j := append(h[:], "Jaguar")
	k := append(h[:], i[1], "Saab", "Nissan")
	k = append(k, k...)
	l := h[:]
	l[0] = "Skoda"
	j[0] = "Bugatti"

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	m := []string{"yellow", "red", "blue"}
	n := m[:]
	fmt.Println(m)
	fmt.Println(n)
	m = append(m, "violet")
	n[0] = "green"
	fmt.Println(m)
	fmt.Println(n)
}
