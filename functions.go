package main

import (
	"fmt"
)

func main() {

	xs:= []float64 {2, 4, 5, 3, 3}
	x := 0;

	result, state := average(xs)
	fmt.Println(result, state)
	fmt.Println(sum(xs...))

	add := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(add(xs[1], xs[3]))

	inc := func() int {
		x++
		return x
	}
	inc()
	fmt.Println(inc())

	nextEven := evenGenerator()
	fmt.Print(nextEven())
	fmt.Print(" ", nextEven())
	fmt.Print(" ", nextEven())
	fmt.Print(" ", nextEven())
	fmt.Println(" ", nextEven())

	fmt.Println(factorial(5))

}

func average(arr []float64) (result float64, error bool) {
	var total float64 = 0
	error = false
	for _, value := range arr {
		total += value
	}
	result = total / float64(len(arr))
	error = true
	return
}

func sum(arr ...float64) (total float64) {
	total = 0
	for _, value := range arr{
		total += value
	}
	return
}

func evenGenerator() func() uint {
	i := uint(0)
	return func() (res uint) {
		res = i
		i += 2
		return
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x*factorial(x - 1)
}