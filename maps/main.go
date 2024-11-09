package main

import "fmt"

func main() {

	grades := make(map[string]int)

	grades["alice"] = 23
	grades["bob"] = 33
	grades["cat"] = 43
	grades["dev"] = 53

	fmt.Println("bob= ", grades["bob"])
	fmt.Println("before bob delete----")

	g1, exist1 := grades["bob"]

	fmt.Println("bob grade= ", g1, " exist=", exist1)
	delete(grades, "bob")
	g2, exist2 := grades["bob"]

	fmt.Println("bob grade= ", g2, " exist=", exist2)

	fmt.Println("loop all key value ---->")

	for ind, val := range grades {
		fmt.Println("index = ", ind, ",value = ", val)
	}
	fmt.Println("\n  initial  during declaratin loop all key value ---->")

	person := map[string]string{
		"aman":  "23",
		"virat": "21",
	}

	for ind, val := range person {
		fmt.Println("index = ", ind, ",value = ", val)
	}

}
