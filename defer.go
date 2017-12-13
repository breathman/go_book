package main

import "fmt"

func main() {
	defer first()
	second()
	defer func() {
		str := recover()
		fmt.Println(str)
	} ()
	panic("Error")
}

func first() {
	fmt.Println("1 function")
}

func second() {
	fmt.Println("2 function")
}
