# GIGMSN Publisher [![Build Status](https://travis-ci.org/gigmsn/publisher.svg?branch=master)](https://travis-ci.org/gigmsn/publisher)

Implement a simple TCP server written in GO for receiving messages though websocket connection and sending them to RabbitMQ message broker.

###Project Structure:

```
.
├── LICENSE
├── Makefile
├── README.md
├── client
│   └── client.js
├── docker-compose.yml
├── main.go
├── resources
│   ├── client
│   │   └── Dockerfile
│   ├── server
│   │   └── prod
│   │       └── Dockerfile
│   └── test
│       └── Dockerfile
└── server
    ├── broker.go
    ├── handlers.go
    ├── handlers_test.go
    ├── server.go
    └── server_test.go
```

### HOWTO:

1. Run unit tests container:

		make test

2. Start **RabbitMQ** message broker:

		make broker/up

	Access UI for managing and monitoring RabbitMQ server at `http://localhost:8080`.

  Default user/password: `guest/guest`

3. Start **server** for serving TCP connections through WebSocket protocol:

		make server/up

4. Start **JS client** for stream messages to the server via websocket connection:

		make client/up

### Available Commands:

```
------------------------------------------------------------------------
GIGMSN Publisher
------------------------------------------------------------------------
broker/stop                    stop and remove broker container
broker/up                      start broker container
client/stop                    stop and remove client container
client/up                      run client container
server/stop                    stop and remove server container
server/up                      run server container
test                           run unit tests
```
