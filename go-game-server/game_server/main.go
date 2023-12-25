package main

import (
	"encoding/json"
	"fmt"
	"gameserver/types"
	"log"
	"math"
	"math/rand"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type PlayerSession struct {
	sessionID int
	clientID  int
	username  string
	inLobby   bool
	conn      *websocket.Conn
}

func newPlayerSession(sid int, conn *websocket.Conn) actor.Producer {
	return func() actor.Receiver {
		return &PlayerSession{
			conn:      conn,
			sessionID: sid,
		}
	}
}

func (s *PlayerSession) Receive(c *actor.Context) {
	switch c.Message().(type){
	case actor.Started:
		s.readLoop()
	}
}

func (s *PlayerSession) readLoop(){
	var msg types.WSMessage
	for{
		if err:= s.conn.WriteJSON(&msg); err !=nil{
			fmt.Println("read error")
			return
		}
		go s.handleMessage(msg)
	}
}

func (s *PlayerSession) handleMessage(msg types.WSMessage){
	switch msg.Type{
	case "login":
		var loginMsg types.Login
		if err := json.Unmarshal(msg.Data, &loginMsg); err!=nil{
			panic(err)
		}
	}
}

type GameServer struct {
	ctx *actor.Context
	sessions map[*actor.PID] struct{}
}

// In actor model, we have actor.receiver and Receive function

func newGameServer() actor.Receiver {
	return &GameServer{
		sessions: make(map[*actor.PID]struct{}),
	}
}

func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		s.ctx = c
		fmt.Println("actor started")
		_ = msg
	}
}

func (s *GameServer) startHTTP() {
	go func() {
		http.HandleFunc("/ws", s.handleWS)
		http.ListenAndServe(":40000", nil)
	}()
}

// handle websocket
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ws upgrade err:", err)
	}
	fmt.Print("new client trying to connect\n")
	sid := rand.Intn(math.MaxInt)
	pid := s.ctx.SpawnChild(newPlayerSession(sid, conn), fmt.Sprintf("session_%d", sid))
	fmt.Printf("client with sid %d and pid %s just connected\n", sid, pid)
	s.sessions[pid] = struct{}{}
}

func main() {
	e, _ := actor.NewEngine()
	e.Spawn(newGameServer, "server")
	select {}
}

