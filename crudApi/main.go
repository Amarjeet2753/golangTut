package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("err while fetching ")
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("res is not getting")
	}

	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("err while getting body")
	// 	return
	// }

	// fmt.Println("data : ", string(data))

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println("err while getting body")
		return
	}

	fmt.Println("todo = ", todo)
	fmt.Println("todo title = ", todo.Title)
	fmt.Println("todo completed = ", todo.Completed)

}

func postRequest() {

	var data = Todo{
		UserId:    1,
		Id:        22222,
		Title:     "heeee;ll",
		Completed: false,
	}
	jsonData, _ := json.Marshal(data)

	jsonstring := string(jsonData)

	jsonReader := strings.NewReader(jsonstring)
	url := "https://jsonplaceholder.typicode.com/todos"
	res, _ := http.Post(url, "application/json", jsonReader)

	defer res.Body.Close()

	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("res data : ", string(resdata))
}

func updateRequest() {
	var data = Todo{
		UserId: 1,

		Title:     "heeee;ll",
		Completed: false,
	}
	jsonData, _ := json.Marshal(data)

	jsonstring := string(jsonData)

	jsonReader := strings.NewReader(jsonstring)
	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, url, jsonReader)

	// Set the content-type header
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer res.Body.Close()

	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("update res data : ", string(resdata))
}

func deleteReq() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	// Create a new HTTP client and send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	fmt.Println("delete  data response : ", res.Status)
}

func main() {

	// postRequest()
	// updateRequest()
	deleteReq()
}
