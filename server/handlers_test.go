package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

// test http status code from root '/' endpoint
func TestIndexHandler(t *testing.T) {
	server := httptest.NewServer(serverEngine())
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("could not make get request to server")
	}

	expectedStatus := 200
	if resp.StatusCode != expectedStatus {
		t.Errorf("expected status code: %d, observed: %d",
			resp.StatusCode, expectedStatus)
	}
}

// test websocket connection on '/ws' endpoint
func TestWsHandler(t *testing.T) {
	server := httptest.NewServer(serverEngine())
	defer server.Close()

	URL, err := url.Parse(server.URL + "/ws")
	if err != nil {
		t.Errorf("could not parse server URL: %s", err)
	}

	URL.Scheme = "ws"

	_, _, err = websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		t.Errorf("expected to create websocket connection with success instead: %s", err)
	}
}
