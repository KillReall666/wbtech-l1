package main

import (
	"fmt"
	"sync"
)

func main() {
	in := []int{2, 4, 6, 8, 10}

	var result int
	var wg sync.WaitGroup
	resultCh := make(chan int)

	for i := 0; i < len(in); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultCh <- n * n
		}(in[i])
	}

	go func() {
		for n := range resultCh {
			result += n
		}
	}()

	wg.Wait()
	close(resultCh)

	fmt.Println(result)
}
