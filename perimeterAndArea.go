package main

import "fmt"

func main() {
	var diameter float32
	diameter = 3.5
	perimeter := diameter * 3.14
	fmt.Println("perimeter is :", perimeter)
	area := 3.14 * (diameter / 2) * (diameter / 2)
	fmt.Println("Area is :", area)
}
