package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)
	toZero(&x)
	fmt.Println(x)

	yPtr := new(int)
	toOne(yPtr)
	fmt.Println(*yPtr)
}

func toZero(xPtr *int) {
	*xPtr = 0
}

func toOne(yPtr *int) {
	*yPtr = 1
}