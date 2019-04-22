package main

import (
	"fmt"
	"github.com/aleksl0l/bomb-backend/frame-runner"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	frameRunner = frame_runner.NewFrameRunner()
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/ws", serveWs)
	addr := viper.GetString("serverAddress")
	log.Info(fmt.Sprintf("Start serving on %s", addr))
	go frameRunner.Run()
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
