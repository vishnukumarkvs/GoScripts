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
    wg := &sync.WaitGroup{}
    // wg.Add(5)

    for i := 0; i < 5; i++ {
		wg.Add(1)
        go fetchUserMatch("User"+strconv.Itoa(i), respch, wg)
    }

	wg.Wait()
	close(respch)

    for user := range respch {
        fmt.Println(user)
    }

	fmt.Println(time.Since(start))
}

func fetchUserMatch(userName string, respch chan string, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 100)
    respch <- userName
	wg.Done() // At last
}
