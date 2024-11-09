package main

import "fmt"

func div(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("error by zero divison")
	}

	return (a / b), nil
}
func div2(a, b float64) (float64, string) {

	if b == 0 {
		return 0, "error by zero divison"
	}

	return (a / b), "nil"
}

func main() {

	// ans, _ := div(10, 4) //egnoe err
	ans, err := div(10, 0) //egnoe err

	if err != nil {
		fmt.Println(err)
	}

	println("ans:", ans)

	ans2, err2 := div(10, 0)

	fmt.Println("ans2:", ans2, "err2", err2)

}
