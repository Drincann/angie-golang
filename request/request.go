package request

import "net/http"

type Request struct {
	req         *http.Request
	Method      string
	Route       string
	queryString map[string][]string
	header      map[string][]string
}

func New(req *http.Request) *Request {
	return &Request{
		req:         req,
		Method:      req.Method,
		Route:       req.URL.Path,
		queryString: req.URL.Query(),
		header:      req.Header,
	}
}

func getFirstFromMapArray(m map[string][]string, key string) string {
	if v, ok := m[key]; ok {
		return v[0]
	} else {
		return ""
	}
}

func getArrayFromMapArray(m map[string][]string, key string) []string {
	if v, ok := m[key]; ok {
		return v
	} else {
		return []string{}
	}
}

func (req *Request) Query(key string) string {
	return getFirstFromMapArray(req.queryString, key)
}

func (req *Request) Querys(key string) []string {
	return getArrayFromMapArray(req.queryString, key)
}

func (req *Request) Header(key string) string {
	return getFirstFromMapArray(req.header, key)
}

func (req *Request) Headers(key string) []string {
	return getArrayFromMapArray(req.header, key)
}

func (req *Request) FromData(key string) string {
	return req.req.FormValue(key)
}
