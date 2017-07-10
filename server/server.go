package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Serve triggers the server initialization
func Serve(addr string) {
	if err := serverEngine().Run(addr); err != nil {
		log.Fatalf("could not serve on %s: %s", addr, err)
	}
}

func serverEngine() *gin.Engine {
	eng := gin.Default()
	// Register resource handlers
	eng.GET("/", indexHandler)
	eng.GET("/ws", wsHandler)
	return eng
}
