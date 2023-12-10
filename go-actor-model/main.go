package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type User struct {
	state uint
}

func newHandler() actor.Receiver{
	return &User{}
}

func (h *User) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Initialized:
		h.state=10
		fmt.Println("Handler initiaized, my state: ", h.state)
	case actor.Started:
		fmt.Println("Handler started")
	case actor.Stopped:
		_ = msg
	}
}

func main() {
	e,_:= actor.NewEngine()
	pid:=e.Spawn(newHandler, "handler")
	fmt.Println("pid->",pid)
}