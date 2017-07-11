package server

import (
	"flag"
	"net"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	flag.StringVar(&queueAddr, "addr", "amqp://guest:guest@broker:5672", "queue address")
	flag.StringVar(&queueName, "queue", "gigmsn_test", "queue name")
	flag.Parse()
}
func TestServe(t *testing.T) {

	timeOut := time.Duration(3) * time.Second
	server := httptest.NewServer(serverEngine())
	defer server.Close()

	// fixes weird double ':' problem
	port := server.URL[len(server.URL)-5:]

	_, err := net.DialTimeout("tcp", "localhost:"+port, timeOut)
	if err != nil {
		t.Errorf("failed to dial server: %s", err)
	}
}

// test if serverEngine returns expected object
func TestServerEngine(t *testing.T) {
	expected := reflect.TypeOf(gin.Default())
	observed := serverEngine()
	if expected != reflect.TypeOf(observed) {
		t.Errorf("observed: %v, expected: %v",
			reflect.TypeOf(observed), expected)
	}
}
