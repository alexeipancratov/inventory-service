package product

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

type Message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket(ws *websocket.Conn) {
	done := make(chan struct{})
	fmt.Println("new websocket connection established")
	go func(c *websocket.Conn) {
		for {
			var msg Message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("received message: %s\n", msg.Data)
		}
		close(done)
	}(ws)

	/*
		In a truly event-based system we would listen for events when some
		product data gets changed and notify the client. Here we're just
		sending new data each 10 seconds.
	*/
loop:
	for {
		select {
		case <-done:
			fmt.Println("connection was closed")
			break loop
		default:
			products, err := getTopTenProducts()
			if err != nil {
				log.Println(err)
				break
			}
			if err := websocket.JSON.Send(ws, products); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Println("closing the websocket")
	defer ws.Close()
}
