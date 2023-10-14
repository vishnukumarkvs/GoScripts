package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	respch := make(chan string, 5)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fetchUserMatch("User"+strconv.Itoa(i), respch, &wg)
	}

	go func() {
		wg.Wait()
		close(respch) // Close the channel when all goroutines are done
	}()

	for user := range respch {
		fmt.Println(user)
	}

	fmt.Println(time.Since(start))
}

func fetchUserMatch(userName string, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	respch <- userName
}
