// In Go version < 1.18 it was not easy to make variant
// of such function for multiple data types.

package main

import "fmt"

func compareInts(x, y int) bool {
	return x < y
}

func compareFloats(x, y float64) bool {
	return x < y
}

func compareStrings(x, y string) bool {
	return x < y
}

func main() {
	fmt.Println(compareInts(1, 2))
	fmt.Println(compareFloats(1.5, 2.6))
	fmt.Println(compareStrings("foo", "bar"))
}
