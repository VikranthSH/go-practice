package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1)
	fmt.Println(cap(s1))
	s1 = append(s1, 60, 70)
	fmt.Println(s1)
	s3 := s1[1:3]
	fmt.Println(s3)
	for i, v := range s1 {

		fmt.Println(i, v)
	}
}
