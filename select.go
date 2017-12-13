package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go one(c1)
	go two(c2)

	go func() {
		for {
			select {
			case msg1:= <-c1 :
				fmt.Println(msg1)
			case msg2:= <-c2:
				fmt.Println(msg2)
			case <- time.After(time.Millisecond*1):
				fmt.Println("timeout")
			default:
				//fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

}

func one(c chan<- string) {
	for i:=0;;i++ {
		c<-"from1"
		time.Sleep(time.Second*10)
	}
}

func two(c chan<- string) {
	for i:=0; ;i++  {
		c<-"from2"
		time.Sleep(time.Second*20)
	}
}