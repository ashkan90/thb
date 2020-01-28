package system

import "net/http"

type Response struct {
	rw http.ResponseWriter
}

func (r *Response) WriteString(d string) {
	_, err := r.rw.Write([]byte(d))
	if err != nil {
		panic(err)
	}
}

func (r *Response) Header() http.Header {
	return r.rw.Header()
}
