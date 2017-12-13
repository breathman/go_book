package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=1; i<=10; i++  {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("Run exit")
}

func f(n int) {
	for i:=1; i<=n; i++  {
		fmt.Println(i, " of ", n)
		time.Sleep(time.Millisecond * 2000)
	}
}
