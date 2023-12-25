package types

type WSMessage struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type Login struct {
	ClientId int    `json:"clientId"`
	Username string `json:"username"`
}