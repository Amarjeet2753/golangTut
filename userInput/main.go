package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// println("enter name :")
	// fmt.Scan(&name)
	// println(" name :", name)

	println("enter name :")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println(" name :", name)
}
