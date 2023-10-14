package main

// https://youtu.be/LGVRPFZr548?si=nI_7xO7ptEOAvgti

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	respch := make(chan any, 2) //buffered channel
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait()     // without this, it wont wait and closes ch immediately
	close(respch) // without this will raise deadlock

	for resp := range respch {
		fmt.Println(resp)
	}

	// fmt.Printf("User Info:-\nName: %+v, Likes: %+v, Match: %+v\n",userName,likes, match)

	fmt.Println(time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Vishnu"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "Ana"
	wg.Done()
}
