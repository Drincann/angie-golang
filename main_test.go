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

type StringReader string

func (str StringReader) Read(p []byte) (n int, err error) {
	return copy(p, []byte(str)), nil
}

func init() {
	app := angie.New()
	go app.Post("/", func(ctx *angie.Context) {
		ctx.SetString("Hello World!")
	}).Get("/", func(ctx *angie.Context) {
		ctx.SetString("Hello World!")
	}).Put("/", func(ctx *angie.Context) {
		ctx.SetString("Hello World!")
	}).Delete("/", func(ctx *angie.Context) {
		ctx.SetString("Hello World!")
	}).Listen(port)
}

func TestGet(t *testing.T) {

	res, err := http.Get(fmt.Sprintf("http://localhost:%d", port))

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

func TestPost(t *testing.T) {

	res, err := http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", StringReader("{}"))

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

func TestPut(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d", port), StringReader("{}"))
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
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

func TestDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:%d", port), StringReader("{}"))
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
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
