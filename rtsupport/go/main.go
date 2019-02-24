package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrater = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4040", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello from go")

	socket, err := upgrater.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}

	}

}
