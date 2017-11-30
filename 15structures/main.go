package main

import "github.com/razorfink/golangbot/15structures/computer"
import "fmt"

type name struct {
	firstName string
	lastName  string
}

func main() {

	// Exported Struct
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "Mac Mini" // cannot access a non-exported field
	fmt.Println("Spec:", spec)

	// Struct Equality
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	fmt.Println("name1", name1)
	fmt.Println("name2", name2)
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	fmt.Println("name3", name3)
	fmt.Println("name4", name4)
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// map types are not comparable
	// so structs containing maps are not comparable
	type image struct {
		data map[int]int
	}

	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	// invalid...not comparable
	// if image1 == image2 {
	// 	fmt.Println("image1 and image2 are equal")
	// }
}
