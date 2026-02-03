package main

import "fmt"

var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
