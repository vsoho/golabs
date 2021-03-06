package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

var i, t int = 6, 9
var a = 8

func main1() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	go say("world")

	fmt.Println("Hello, world")
	fmt.Println("hello ", rand.Intn(10))

	for i := 0; i < 8; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("comp")
	}

}
