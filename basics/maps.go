package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["India"] = 1
	m["USA"] = 2
	m["Australia"] = 3
	fmt.Println(m)

	k := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(k)
}
