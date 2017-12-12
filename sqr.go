package main

import "fmt"

var input, output float64

func main() {

	fmt.Print("Enter a number ")
	fmt.Scanf("%f", &input)
	output = input * input
	fmt.Println("Result: ", output)
}
