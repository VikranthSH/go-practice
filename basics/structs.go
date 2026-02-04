package main

import "fmt"

// type car struct{
// 	wheels int
// 	name string
// 	model int
// 	FrontWheel wheel
// 	BackWheel wheel
// }
// type wheel struct{
// 	Radius int
// 	Material string
// }
// func main(){
// 	mycar:=car{}
// 	mycar.wheels=4
// 	mycar.FrontWheel.Radius=7
// 	fmt.Println(mycar.wheels)

// }

type rect struct{
	length int
	breadth int
}
func(r rect) area() int{
	return r.length*r.breadth
}
func main(){
r:=rect{
	length:5,
	breadth:7,
}
fmt.Println(r.area())}