package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	// router
	go app.Post("/router/post", func(ctx *angie.Context) {
		ctx.ResString("Hello World!")
	}).Get("/router/get", func(ctx *angie.Context) {
		ctx.ResString("Hello World!")
	}).Put("/router/put", func(ctx *angie.Context) {
		ctx.ResString("Hello World!")
	}).Delete("/router/delete", func(ctx *angie.Context) {
		ctx.ResString("Hello World!")
	}) /*.Get("/router/no-match", func(ctx *angie.Context) {})*/

	// webContext
	type j map[string]interface{}
	app.Get("/webContext/resjson", func(ctx *angie.Context) {
		ctx.ResJson(j{
			"hello": "world",
			"obj": j{
				"hello": "world",
			},
		})
	}).Get("/webContext/resbytes", func(ctx *angie.Context) {
		ctx.ResBytes([]byte("Hello World!"))
	}).Get("/webContext/resstring", func(ctx *angie.Context) {
		ctx.ResString("Hello World!")
	}).Get("/webContext/setstatus", func(ctx *angie.Context) {
		ctx.SetStatus(http.StatusNotFound)
	}).Get("/webContext/setheader", func(ctx *angie.Context) {
		ctx.SetHeader("hello", "world")
	})

	go app.Listen(port)
}

func TestNoMatchRoute(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/router/no-match", port))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, res.StatusCode)
	}

}

func TestResJson(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/webContext/resjson", port))

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

	// check header
	if res.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected header 'Content-Type: application/json', got '%s'", res.Header.Get("Content-Type"))
	}

	// regexp body
	if !regexp.MustCompile(`{"hello":"world","obj":{"hello":"world"}}`).Match(body) {
		t.Errorf("Expected body '{\"hello\":\"world\",\"obj\":{\"hello\":\"world\"}}', got '%s'", string(body))
	}
}

func TestResBytes(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/webContext/resbytes", port))

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

	// check header
	if res.Header.Get("Content-Type") != "application/octet-stream" {
		t.Errorf("Expected header 'Content-Type: application/octet-stream', got '%s'", res.Header.Get("Content-Type"))
	}

	if string(body) != "Hello World!" {
		t.Errorf("Expected body 'Hello World!', got '%s'", string(body))
	}

}

func TestResString(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/webContext/resstring", port))

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

	// check header
	if res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Expected header 'Content-Type: text/plain; charset=utf-8', got '%s'", res.Header.Get("Content-Type"))
	}

	if string(body) != "Hello World!" {
		t.Errorf("Expected body 'Hello World!', got '%s'", string(body))
	}
}

func TestSetStatus(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/webContext/setstatus", port))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestSetHeader(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d/webContext/setheader", port))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// read response
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	// check header
	if res.Header.Get("hello") != "world" {
		t.Errorf("Expected header 'hello: world', got '%s'", res.Header.Get("hello"))
	}
}

func TestGet(t *testing.T) {

	res, err := http.Get(fmt.Sprintf("http://localhost:%d/router/get", port))

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

	res, err := http.Post(fmt.Sprintf("http://localhost:%d/router/post", port), "application/json", StringReader("{}"))

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
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/router/put", port), StringReader("{}"))
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
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:%d/router/delete", port), StringReader("{}"))
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
