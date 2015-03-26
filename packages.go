package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("My favourite number is ", rand.Intn(10))
}
