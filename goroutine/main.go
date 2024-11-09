package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker started : ", i)
	fmt.Println("worker ended : ", i)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
