package player

import (
	"github.com/aleksl0l/bomb-backend/game-object"
	"github.com/aleksl0l/bomb-backend/message"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type IPlayer interface {
	GetAction() *game_object.GameObject
	SendMessage(data []byte) error
}

type Player struct {
	Player *message.Player
	conn   *websocket.Conn
}

func NewPlayer(conn *websocket.Conn, player *message.Player) *Player {
	return &Player{
		player,
		conn,
	}
}

func (p *Player) HandlePlayerActions() {
	defer func() {
		p.conn.Close()
	}()
	p.conn.SetReadLimit(maxMessageSize)
	p.conn.SetReadDeadline(time.Now().Add(pongWait))
	p.conn.SetPongHandler(func(string) error {
		p.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, msg, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		parsedMsg := &message.Message{}
		err = proto.Unmarshal(msg, parsedMsg)
		if err != nil {
			continue
		}
		switch parsedMsg.Type {
		case message.MessageType_PLAYER_POSITION_REQUEST:
			position := &message.PlayerPositionRequest{}
			err = proto.Unmarshal(parsedMsg.Content, position)
			p.Player.Position = position.Position
			//case message.MessageType_CREATE_GAME_REQUEST:

		}
	}
}

func (p *Player) GetAction() *game_object.GameObject {
	return nil
}

func (p *Player) SendMessage(data []byte) error {
	return p.conn.WriteMessage(websocket.TextMessage, data)
}
