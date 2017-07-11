package server

import (
	"net/http"

	"github.com/gigmsn/messagebroker"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "GIGMSN")
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("could not upgrade connection to use websocket: %s", err)
	}
	go ws(conn)
}

func ws(conn *websocket.Conn) {
	broker, err := messagebroker.New(queueAddr, queueName)
	if err != nil {
		log.Fatalf("could not create broker: %s", err)
	}
	defer broker.Close()
	log.Infoln("broker connection established")

	msgCh := make(chan []byte, 100)
	doneCh := make(chan bool)

	go broker.Publish(msgCh, doneCh)

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Infoln("connection closed")

			// Close connection
			conn.Close()

			// Close amqp channel
			broker.Channel.Close()

			// notify publisher to close msg channel
			<-doneCh
			return
		}
		conn.WriteMessage(1, []byte("acknowledged"))
		msgCh <- msg
	}
}
