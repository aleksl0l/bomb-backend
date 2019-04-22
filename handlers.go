package main

import (
	"fmt"
	"github.com/aleksl0l/bomb-backend/message"
	"github.com/aleksl0l/bomb-backend/player"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		return
	}
	msg := &message.Message{}
	err = proto.Unmarshal(data, msg)
	if err != nil {
		return
	}
	switch msg.Type {
	case message.MessageType_JOIN_GAME:
		err = handleJoinGame(msg, conn)
		if err != nil {
			return
		}
	case message.MessageType_LEAVE_GAME:
		err = handleLeaveGame(msg, conn)
		if err != nil {
			return
		}
	}
	log.Info("read done")
}

func handleJoinGame(msg *message.Message, conn *websocket.Conn) error {
	joinGame := &message.JoinGame{}
	err := proto.Unmarshal(msg.Content, joinGame)
	if err != nil {
		return err
	}
	newPlayer := player.NewPlayer(conn)
	frameRunner.AddPlayer(newPlayer)
	log.Info(fmt.Sprintf("User %s joined game", joinGame.Username))
	return nil
}

func handleLeaveGame(msg *message.Message, conn *websocket.Conn) error {
	return nil
}
