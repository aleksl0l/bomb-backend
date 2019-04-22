package player

import (
	"github.com/aleksl0l/bomb-backend/game-object"
	"github.com/aleksl0l/bomb-backend/message"
	"github.com/gorilla/websocket"
)

type Connection struct {
	conn websocket.Conn
}

func (c *Connection) Write(data []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Connection) Read() (error, []byte) {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		return err, nil
	}
	return nil, msg
}

type IPlayer interface {
	GetAction() *game_object.GameObject
	SendMessage(data []byte) error
}

type Player struct {
	Player message.Player
	conn   *websocket.Conn
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		message.Player{},
		conn,
	}
}

func (p *Player) GetAction() *game_object.GameObject {
	return nil
}

func (p *Player) SendMessage(data []byte) error {
	return p.conn.WriteMessage(websocket.TextMessage, data)
}
