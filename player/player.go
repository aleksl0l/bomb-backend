package player

import (
	"github.com/aleksl0l/bomb-backend/game-map"
	"github.com/aleksl0l/bomb-backend/game-object"
	"github.com/gorilla/websocket"
)

type Connection struct {
	conn websocket.Conn
}

type PlayerPosition interface {
}

func (c *Connection) Write(data []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Connection) Read() (error, []byte) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return err, nil
	}
	return nil, message
}

type IPlayer interface {
	GetAction() *game_object.GameObject
	SendMap(gameMap game_map.GameMap)
}

type Profile struct {
	username string
	points   rune
}

type Player struct {
	conn    Connection
	profile Profile
}

func NewPlayer(conn Connection, profile Profile) *Player {
	return &Player{
		conn,
		profile,
	}
}

func (p *Player) GetAction() *game_object.GameObject {
	return nil
}

func (p *Player) SendMap(gameMap game_map.GameMap) {
	return
}
