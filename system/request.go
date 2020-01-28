package system

import (
	"net/http"
	"net/url"
)

type Request struct {
	request *http.Request
}

func (r *Request) All() url.Values {
	return r.request.Form
}

func (r *Request) Except(keys []string) url.Values {
	values := make(url.Values)
	for formKey, formVal := range r.request.Form {
		for _, key := range keys {
			if formKey != key {
				values[formKey] = formVal
			}
		}

	}

	return values
}

func (r *Request) Get(key string) []string {
	return r.All()[key]
}

func (r *Request) Has(key string) bool {
	_, ok := r.All()[key]
	return ok
}

func GetRequest() *Request {
	return GetApplication().req
}
