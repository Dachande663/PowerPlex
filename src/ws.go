package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type wsMessage struct {
	Action string `json:"action"`
	Source string `json:"source"`
	Data   string `json:"data"`
}

func startWebsocket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		println("shit1")
		http.Error(w, "Unable to connect to system", 500)
		return
	}

	msgs := make(chan wsMessage)

	defer func() {
		ws.Close()
		close(msgs)
	}()

	go wsWriter(ws, msgs)
	wsReader(ws, msgs)

}

func wsWriter(ws *websocket.Conn, msgs chan wsMessage) {

	for {
		msg := <-msgs
		ws.WriteJSON(msg)
	}

}

func wsReader(ws *websocket.Conn, msgs chan wsMessage) {

	for {

		msg := wsMessage{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			println("ReadJSON err:", err.Error())
			return
		}

		switch msg.Action {
		case "export":
			go wsExportCommand(msgs)
		default:
			msg.Source = "server"
			msgs <- msg
		}

	}

}

func wsExportCommand(output chan wsMessage) {

	output <- wsMessage{Action: "log", Source: "server", Data: "start"}

	exporter := makeExporter()

	exporter.Reporter = func(msg string) {
		output <- wsMessage{Action: "log", Source: "server", Data: msg}
	}

	library := exporter.Export()
	j, _ := json.MarshalIndent(library, "", "  ")

	output <- wsMessage{Action: "log", Source: "server", Data: "end"}

	cwd, _ := os.Getwd()
	dir, _ := filepath.Abs(cwd)
	file := dir + "/data.json"

	println(file)

	ioutil.WriteFile(file, j, 0644)

	output <- wsMessage{Action: "log", Source: "server", Data: file}

}
