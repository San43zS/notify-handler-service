package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("websocket")

func New() (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/ws", nil)
	if conn != nil {
		log.Infof("websocket connection established")
	}

	if err != nil {
		log.Criticalf("websocket connection error: %v", err)
		return nil, fmt.Errorf("websocket connection error: %v", err)
	}

	return conn, nil
}
