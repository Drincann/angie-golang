package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Drincann/angie-golang/angie"
)

const (
	port = 8080
)

func TestBlack(t *testing.T) {
	app := angie.New()
	go app.Get("/", func(ctx *angie.Context) {
		ctx.Res.Body = "Hello World!"
	}).Listen(port)

	res, err := http.Get(fmt.Sprintf("http://localhost:%d/?name=Drincann", port))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// read response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(body) != "Hello World!" {
		t.Errorf("Expected body 'Hello World!', got '%s'", string(body))
	}
}
