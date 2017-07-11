package main

import (
	"flag"

	"github.com/gigmsn/publisher/server"
)

func main() {
	addrPtr := flag.String("addr", "amqp://guest:guest@broker:5672", "queue address")
	queuePrt := flag.String("queue", "gigmsn", "queue name")
	flag.Parse()

	server.Serve(":3000", *addrPtr, *queuePrt)
}
