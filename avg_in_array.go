package main

import "fmt"

func main() {
	x:= [5]float64 {2, 4, 5, 3, 3}
	var total float64 = 0;
	for _, value := range x {
		total += value
	}
	result := total / float64(len(x))
	fmt.Print(result)
}
