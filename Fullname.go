package main

import "fmt"

func main() {
	fmt.Println("Enter first name : ")
	var (
		first string
		last  string
	)
	fmt.Scan(&first)
	fmt.Println("Enter last name : ")
	fmt.Scan(&last)
	fmt.Println(first, last)
}
