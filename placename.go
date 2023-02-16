package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the name of place : ")
	var placename string
	fmt.Scan(&placename)
	capplace := strings.ToUpper(placename)
	fmt.Println(capplace)
	fmt.Println(placename)

}
