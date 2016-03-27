package main

import (
	"fmt"
	"math/rand"
	"time"
)

var i = 6

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	go say("world")

	fmt.Println("Hello, world")
	fmt.Println("hello ", rand.Intn(10))

	for i := 0; i < 8; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("comp")
	}
}
