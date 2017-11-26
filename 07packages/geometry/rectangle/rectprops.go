package rectangle

import "math"
import "fmt"

func init() {
	fmt.Println("rectangle package initialized")
}

// Area returns the rectangular area for the given dimensions
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal returns the rectangular diagonal for the given dimensions
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
