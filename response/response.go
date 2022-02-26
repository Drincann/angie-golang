package response

import "net/http"

type Response struct {
	Body   string
	Status int
}

func New(res http.ResponseWriter) *Response {
	return &Response{
		Status: http.StatusOK,
		Body:   "",
	}
}
