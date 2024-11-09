package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Enter num to check odd even : ")
	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')

	num = strings.TrimSpace(num)

	// n, _ := strconv.Atoi(num)

	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	// fmt.Printf("type of n = %T\n", n)
	fmt.Println("n=", n)

	if n%2 == 0 {
		fmt.Println("n is even")
	} else {
		fmt.Println("n is odd")

	}
	fmt.Println("\n switch -----------")

	day := 56

	switch day {
	case 1, 56:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	default:
		fmt.Println("other")
	}
	fmt.Println("\n for loop -----------")

	for i := 0; i < 5; i++ {

		fmt.Println("i=", i)
	}

	nums := []int{23, 45, 6, 7}

	for ind, val := range nums {
		println("index : ", ind, "val=", val)
	}

}
