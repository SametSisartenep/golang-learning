package main

import "fmt"

func main() {
	var x [10]int

	fmt.Println("Give me ten numbers, separated by whitespaces:")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Scanf("%d", &x[i])
		fmt.Printf("Just got ==> x[%d] : %d\n", i, x[i])
	}

	fmt.Println("Those are the selected numbers: ")

	for i := 0; i < 10; i++ {
		fmt.Println(x[i])
	}

	fmt.Println("And the minimum of all those is: ")

	var minimum int

	for i := 0; i < 10; i++ {
		if i == 0 {
			minimum = x[i]
		} else {
			if minimum > x[i] {
				minimum = x[i]
			}
		}
	}

	fmt.Printf("%d", minimum)
}
