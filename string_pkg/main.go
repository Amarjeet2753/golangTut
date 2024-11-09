package main

import (
	"fmt"
	"strings"
)

func main() {

	var names string = "rohit, virat,hardik,chahal"

	parts := strings.Split(names, ",")

	fmt.Println("names=", parts)

	var str string = "     hello world   "

	trim := strings.TrimSpace(str)

	fmt.Println("trim=", trim)

	s1 := "amar"
	s2 := "prajapati"

	res := strings.Join([]string{s1, s2}, ",,")
	fmt.Println("res=", res)

}
