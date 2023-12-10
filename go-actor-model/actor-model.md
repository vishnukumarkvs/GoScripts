```
package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type User struct {
	state uint
}

type SetState struct{
	value uint
}

type ResetState struct{}

func newHandler() actor.Receiver{
	return &User{}
}

func (h *User) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case SetState:
		h.state = msg.value
		fmt.Println("handler received new state:", h.state)
	case ResetState:
		h.state=0
		fmt.Println("Handler reset state: ",h.state)
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
	for i:=0;i<10;i++{
		go e.Send(pid,SetState{value: uint(i)})
	}
	e.Send(pid, ResetState{}) // msg has type any
	// we are sending msg to pid
}
```
