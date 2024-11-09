package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func mul(a int, b int) (result int) {

	result = a * b
	return
}

// func person(name, age) {
func person(name string, age int) {
	fmt.Println("name: ", name, "age: ", age)
}
func main() {

	sum := add(3, 4)
	fmt.Println("sum=", sum)
	fmt.Println("mul=", mul(4, 5))

	person("virat", 45)

}
