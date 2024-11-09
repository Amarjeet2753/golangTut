package main

import (
	"fmt"
	"strconv"
)

func change(ptr *int) {
	*ptr = *ptr * 2
}

func main() {

	var num int = 5

	// var ptr *int

	ptr := &num

	fmt.Println("num=", num)
	fmt.Println("&num=", ptr)
	fmt.Println("num vale0=", *ptr)

	n1 := 8
	// ptr1:=&s;

	change(&n1)
	fmt.Println("n1=", n1)

	fmt.Println("data conversion =------------->")

	str := "123"

	ns, _ := strconv.Atoi(str)

	fmt.Println("converte ns =", ns)
	fmt.Printf("type of ns %T", ns)

	f1 := "13"

	fc, _ := strconv.ParseFloat(f1, 64)
	fmt.Printf("type of fc %T", fc)

}
