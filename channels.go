package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}

func pinger(c chan<- string) {
	for i:=0;;i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i:=0;;i++ {
		c <- "pong"
	}
}

func printer(c <-chan string)  {
	for i:=0; ;i++  {
		msg:= <- c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}