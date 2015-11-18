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

type wsMessage struct {
	Code   int    `json:"code"`
	Action string `json:"action"`
	Data   string `json:"data"`
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func startWebsocket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
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

// Handle sending messages to the front-end via websockets
func wsWriter(ws *websocket.Conn, msgs chan wsMessage) {

	for {
		msg := <-msgs
		ws.WriteJSON(msg)
	}

}

// Handle reading messages sent over websockets from the front-end
func wsReader(ws *websocket.Conn, msgs chan wsMessage) {

	for {

		msg := wsMessage{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			return // abort reader
		}

		switch msg.Action {
		case "export":
			go wsExportCommand(msgs)
		default:
			msgs <- msg
		}

	}

}

// Handle front-end requesting an export
func wsExportCommand(output chan wsMessage) {

	output <- wsMessage{Code: 200, Action: "log", Data: "start"}

	exporter := makeExporter()

	exporter.Reporter = func(msg string) {
		output <- wsMessage{Code: 200, Action: "log", Data: msg}
	}

	library := exporter.Export()
	j, _ := json.MarshalIndent(library, "", "  ")

	output <- wsMessage{Code: 200, Action: "log", Data: "end"}

	cwd, _ := os.Getwd()
	dir, _ := filepath.Abs(cwd)
	file := dir + "/data.json"

	ioutil.WriteFile(file, j, 0644)

	output <- wsMessage{Code: 200, Action: "log", Data: file}

}
