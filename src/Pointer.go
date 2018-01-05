package main

import "fmt"

/*func zero(x int){
	x = 0
}
*/

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(&x)
	fmt.Println("Get address of variable x in memory===> ", &x)
	fmt.Println("Get value of xPtr in memory===> ", x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	xx := 1.5
	square(&xx)
	fmt.Println("The value of xx===> ", xx)

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("Value of x ==>", a, " Value of y ==>", b)
}

func square(x *float64) {
	//*x = *x * *x
	*x = *x * *x
}

func swap(x *int, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}
