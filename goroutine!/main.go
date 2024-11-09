package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello")
}
func sayHello2() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("hello2")
}

func main() {

	go sayHello()
	go sayHello2()

	time.Sleep(2000 * time.Millisecond)
}
