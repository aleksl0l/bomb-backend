package frame_runner

import (
	"github.com/aleksl0l/bomb-backend/game-map"
	"github.com/aleksl0l/bomb-backend/player"
	"github.com/gorilla/websocket"
)

type IFrameRunner interface {
	broadcastMap()
}

type FrameRunner struct {
	Players []player.IPlayer
	Map     game_map.GameMap
	Conn    websocket.Conn
}

func NewFrameRunner() {

}

func (fr *FrameRunner) broadcastMap() {
	for i := 0; i < len(fr.Players); i++ {
		fr.Players[i].SendMap(fr.Map)
	}
}
