package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Is_admin bool   `json:"is_admin"`
}

func main() {

	p1 := Person{
		Name: "Virat",
		Age:  21,
	}

	jsonData, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("erron in marslling")
		return

	}

	fmt.Println("json data = ", string(jsonData))

	// decoding

	var p1data Person

	err = json.Unmarshal(jsonData, &p1data)

	if err != nil {
		fmt.Println("erron in ummarslling")
		return

	}

	fmt.Println("unmarshal data : ", p1data)

}
