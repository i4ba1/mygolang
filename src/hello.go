package main

import "fmt"

func hello() {
	var x string = "Hello GO...."
	fmt.Println(x)

	a := 1
	b := 2
	fmt.Println(a + b)

	fmt.Println("The weather in celcius==> ", fahrenheitToCelcius(100))
	fmt.Println(add(1, 2, 3))
}

func fahrenheitToCelcius(f int) int {
	c := (f - 32) * 5 / 9

	return c
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}
