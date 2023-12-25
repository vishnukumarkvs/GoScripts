package main

import (
	"encoding/json"
	"gameserver/types"
	"log"
	"math"
	"math/rand"

	"github.com/gorilla/websocket"
)


type GameClient struct {
	conn     *websocket.Conn
	clientId int
	username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		conn:     conn,
		clientId: rand.Intn(math.MaxInt),
		username: username,
	}
}

func (c *GameClient) login() error {
	b, err := json.Marshal(types.Login{
		ClientId: c.clientId,
		Username: c.username,
	})
	if err!=nil{
		return err
	}
	msg:= types.WSMessage{
		Type: "login",
		Data: b,
	}
	return c.conn.WriteJSON(msg)
}

const wsServerEndpoint = "ws://localhost:40000/ws"

func main() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal("Not able to dial to server endpoint")
	}

	c := newGameClient(conn, "vkvs")

	if err := c.login(); err != nil {
		log.Fatal("Error: ", err)
	}
}
