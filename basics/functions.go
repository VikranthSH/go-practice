package main

import "fmt"

func add(a int, b int) float64 {
	return float64(a + b)
}
func maths() (int, int) {
	return 5, 6
}
func main() {
	result := add(6, 4)
	fmt.Println("6+4 = ", result)
	a, b := maths()
	fmt.Println("Values from maths function:", a, b)
}
