package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

//Go语言标准包里面没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var str string
		if err = websocket.Message.Receive(ws, &str); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("receive message : ", str)
		msg := "Received: " + str
		fmt.Println("send message : ", msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println(err)
			break
		}
	}
}
func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
