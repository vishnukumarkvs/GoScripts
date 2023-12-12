package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type GameServer struct{}

// In actor model, we have actor.receiver and Receive function

func newGameServer() actor.Receiver{
	return &GameServer{}
}

func (s *GameServer) Receive(c *actor.Context){
	switch msg:= c.Message().(type){
	case actor.Started:
		fmt.Println("actor started")
	default:
		_ = msg
	}
}

func (s *GameServer) startHTTP(){
	go func(){
		http.HandleFunc("/ws",s.handleWS)
		http.ListenAndServe(":40000",nil)
	}()
}

// handle websocket
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Fatal("ws upgrade err:" ,err)
	}
	fmt.Print(conn)
}

func main(){
	e, _ := actor.NewEngine()
	e.Spawn(newGameServer, "server")
}