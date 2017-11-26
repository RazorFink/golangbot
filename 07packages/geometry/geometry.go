package main

import "fmt"
import rect "github.com/razorfink/golangbot/07packages/geometry/rectangle"
import "log"

var rectLen, rectWidth float64 = -6, 7

var _ = rect.Area

func init() {

	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than 0")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than 0")
	}
}

func main() {
	fmt.Println("Geometric shape properties")

	// fmt.Printf("Aread of rectangle: %.2f\n", rectangle.Area(rectLen, rectWidth))
	// fmt.Printf("Diagonal of rectangle: %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
