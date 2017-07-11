# GIGMSN Publisher [![Build Status](https://travis-ci.org/gigmsn/publisher.svg?branch=master)](https://travis-ci.org/gigmsn/publisher)

Implement a simple TCP server written in GO for receiving messages though websocket connection and sending them to RabbitMQ message broker.

### Project Structure:

```
.
├── LICENSE
├── Makefile
├── README.md
├── jsclient
│   └── jsclient.js
├── docker-compose.yml
├── main.go
├── resources
│   ├── jsclient
│   │   └── Dockerfile
│   ├── server
│   │   └── prod
│   │       └── Dockerfile
│   └── test
│       └── Dockerfile
└── server
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

4. Start **jsclient** for stream messages to the server via websocket connection:

		make jsclient/up

### Available Commands:

```
------------------------------------------------------------------------
GIGMSN Publisher
------------------------------------------------------------------------
broker/stop                    stop and remove broker container
broker/up                      start broker container
jsclient/stop                  stop and remove jsclient container
jsclient/up                    run jsclient container
server/stop                    stop and remove server container
server/up                      run server container
test                           run tests
```
