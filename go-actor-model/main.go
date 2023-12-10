package main

type User struct {
	state uint
}

func (h *User) handleFunc(state int) {
	h.state = uint(state)
}

func main() {
	user1 := User{}
	for i := 0; i < 10; i++ {
		go user1.handleFunc(i)
	}
}

// go run --race main.go -> data race occurs