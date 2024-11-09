package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("err while fetching url")
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	fmt.Println("res=", string(data))

	fmt.Println("\n \n URL handling \n")

	myurl := "https://www.youtube.com/watch?v=gVsCscDvKK0&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=27"

	parseUrl, err := url.Parse(myurl)

	fmt.Println("parseUrl=", parseUrl)

	fmt.Println("scheme : ", parseUrl.Scheme)
	fmt.Println("path : ", parseUrl.Path)
	fmt.Println("host : ", parseUrl.Host)
	fmt.Println("rawQuery : ", parseUrl.RawQuery)
}
