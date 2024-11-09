package main

import "fmt"

func main() {

	// Array
	var name [5]string
	name[0] = "virat"
	fmt.Println("name = ", name, "len=", len(name), "cap=", cap(name))
	fmt.Printf("name = %q\n", name)

	// var num = [5]int{1, 2, 3, 4}
	num := [5]int{1, 2, 3, 4}
	fmt.Println("num=", num)
	fmt.Printf("type of num =%T\n \nSlice from here : \n \n", num)

	// Slice-------------------------------------------------->

	n := []int{1, 2, 4}

	n = append(n, 5, 6)
	fmt.Println("n==", n)

	// use make funcion len ,cap

	// n2 := make([]int,0) for empty
	// n2 := []int{} for empty
	n2 := make([]int, 2, 4)

	n2 = append(n2, 2, 3)
	n2 = append(n2, 2)

	fmt.Println("n2 =", n2, "len=", len(n2), "cap=", cap(n2))

}
