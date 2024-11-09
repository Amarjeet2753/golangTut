package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello, World!  shh")

	myutil.PrintMessage("hellow from myutil")

	var name string = "virat"
	var name2 = "rohit" //auto convert

	num := 3
	var f = true
	f1 := false

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(num)
	fmt.Println(f)
	fmt.Println(f1)

	age := 23
	hight := 4.555

	fmt.Println("age=", age, "hight=", hight)
	fmt.Printf("name = %s ,age = %d ,hight = %0.2f \n", name, age, hight)

	fmt.Printf("type of name = %T", name)

}
