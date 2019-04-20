package main

import (
	"fmt"
	"github.com/aleksl0l/bomb-backend/message"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
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
			handleJoinGame(msg)
		case message.MessageType_LEAVE_GAME:
			handlerLeaveGame(msg)
		}
		log.Info("read done")
	}
}

func handleJoinGame(msg *message.Message) error {
	joinGame := &message.JoinGame{}
	err := proto.Unmarshal(msg.Content, joinGame)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("User %s join game", joinGame.Username))
	return nil
}

func handlerLeaveGame(msg *message.Message) error {
	return nil
}

func main() {
	http.HandleFunc("/ws", serveWs)
	addr := viper.GetString("serverAddress")
	log.Info(fmt.Sprintf("Start serving on %s", addr))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
