package response

import (
	"bytes"
	"net/http"
)

type Response struct {
	Body   bytes.Buffer
	Status int
}

func New(res http.ResponseWriter) *Response {
	return &Response{
		Status: http.StatusOK,
	}
}
