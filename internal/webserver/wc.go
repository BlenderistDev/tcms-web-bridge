package webserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// getWcHandler handler for websockets
func getWcHandler(addConsumer chan chan []uint8) func(c *gin.Context) {
	return func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		log.Println(err)

		defer func(ws *websocket.Conn) {
			err := ws.Close()
			log.Println(err)
		}(ws)

		ch := make(chan []uint8)
		addConsumer <- ch

		for {
			data := <-ch
			err = ws.WriteMessage(websocket.TextMessage, data)
			log.Println(err)
		}
	}
}
