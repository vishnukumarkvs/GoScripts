package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	likes := fetchUserLikes(userName)
	match := fetchUserMatch(userName)

	fmt.Printf("User Info:-\nName: %+v, Likes: %+v, Match: %+v\n", userName, likes, match)

	fmt.Println(time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Vishnu"
}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150)

	return 10
}

func fetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 100)

	return "Ana"
}
