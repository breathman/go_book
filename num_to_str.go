package main

import "fmt"

var input int
var output string

func main() {
	fmt.Print("Enter a number ");
	fmt.Scanf("%d", &input)
	switch input {
	case 1: output = "One"
	case 2: output = "Two"
	case 3: output = "Three"
	case 4: output = "Four"
	case 5: output = "Five"
	case 6: output = "Six"
	case 7: output = "Seven"
	case 8: output = "Eight"
	case 9: output = "Nine"
	default: output = "Unknown number"
	}
	fmt.Println(output)
}
