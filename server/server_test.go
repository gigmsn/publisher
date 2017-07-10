package server

import (
	"net"
	"net/http/httptest"
	"testing"
	"time"

	"reflect"

	"github.com/gin-gonic/gin"
)

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
