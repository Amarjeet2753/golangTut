package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {

	// creating file------------------------------------------

	// file, err := os.Create("abc.txt")

	// if err != nil {
	// 	fmt.Println("Error while creating file")

	// }

	// defer file.Close()

	// content := "hello this is from file hanlding code"

	// byte, err := io.WriteString(file, content)

	// fmt.Println("file created successfully size= ", byte)
	// fmt.Println("file created successfully")

	// reading file-------------------------------------->

	// file, err := os.Open("abc.txt")
	// if err != nil {
	// 	fmt.Println("Error while opening file")

	// }

	// defer file.Close()

	// // craete bufer

	// bufer := make([]byte, 1024)

	// for {
	// 	n, err2 := file.Read(bufer)

	// 	if err2 == io.EOF {
	// 		break
	// 	}
	// 	if err2 != nil {
	// 		fmt.Println("Error while reading file")

	// 	}
	// 	fmt.Println("file content : ", string(bufer[:n]))
	// }

	// read file method 2------------------>

	content, _ := os.ReadFile("abc.txt")

	fmt.Println(string(content))

}
