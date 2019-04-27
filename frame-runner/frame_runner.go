package frame_runner

import (
	"fmt"
	"github.com/aleksl0l/bomb-backend/game-map"
	"github.com/aleksl0l/bomb-backend/message"
	"github.com/aleksl0l/bomb-backend/player"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type IFrameRunner interface {
	BroadcastMap()
	AddPlayer(player player.IPlayer)
	Run()
	prepareMapResponse() ([]byte, error)
}

type FrameRunner struct {
	Players []player.IPlayer
	Map     game_map.GameMap
}

func NewFrameRunner() IFrameRunner {
	return &FrameRunner{
		Players: make([]player.IPlayer, 0, 10),
		Map:     game_map.NewNaiveMap(10),
	}
}

func (fr *FrameRunner) prepareMapResponse() ([]byte, error) {
	naiveMap := fr.Map.(*game_map.NaiveMap)
	gameMapResponse := &message.GameMap{
		GameObjects: naiveMap.Objects,
		Players:     make([]*message.Player, 0, len(fr.Players)),
	}
	for _, p := range fr.Players {
		gameMapResponse.Players = append(gameMapResponse.Players, p.(*player.Player).Player)
	}
	content, err := proto.Marshal(gameMapResponse)
	if err != nil {
		return nil, err
	}
	messageResponse := &message.Message{
		Type:    message.MessageType_MAP_RESPONSE,
		Content: content,
	}
	return proto.Marshal(messageResponse)
}

func (fr *FrameRunner) BroadcastMap() {
	log.Info("Start Broadcast Map")
	data, err := fr.prepareMapResponse()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(fr.Players); i++ {
		fr.Players[i].SendMessage(data)
		log.Info(fmt.Sprintf("Broadcast map to %s", fr.Players[i]))
	}
}

func (fr *FrameRunner) AddPlayer(player player.IPlayer) {
	fr.Players = append(fr.Players, player)
}

func (fr *FrameRunner) Run() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fr.BroadcastMap()
		}
	}
}
