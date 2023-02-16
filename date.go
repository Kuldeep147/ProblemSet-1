package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	year := time.Now().Local().Year()
	day := time.Now().Day()
	month := time.Now().Month()
	fmt.Println(year)
	fmt.Println(day)
	fmt.Println(month)
}
