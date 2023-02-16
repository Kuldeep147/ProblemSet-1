package main

import (
	"fmt"
)

func main() {
	var i float32
	fmt.Println("Enter temperature in celsius : ")
	fmt.Scan(&i)
	f := (i * 9 / 5) + 32
	fmt.Println("In Fahrenheit : ", f)
	var x = 5.5
	if x > 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("Less than or equal to 5")
	}
	t := '😀'
	fmt.Printf("data type %c is %T and the rune value %U \n", t, t, t)

	s := "Hello, 日本語"
	fmt.Println(s)
	some := 42
	fmt.Println(some)

	var pointerexample *int = &some
	fmt.Println(*pointerexample)
}
