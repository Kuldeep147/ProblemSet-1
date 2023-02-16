package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a1 := rand.Intn(40) + 10
	a2 := rand.Intn(40) + 10
	a3 := rand.Intn(40) + 10
	a4 := rand.Intn(40) + 10
	a5 := rand.Intn(40) + 10
	fmt.Println(a1, a2, a3, a4, a5)
	avg := (a1 + a2 + a3 + a4 + a5) / 5
	fmt.Println("Average is : ", avg)
}
