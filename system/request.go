package system

import (
	"github.com/go-playground/validator"
	"net/http"
	"net/url"
)

type Method string

const (
	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

type Request struct {
	request *http.Request
}

type IRequest interface {
	All() url.Values
	Except([]string) url.Values
	Get(string) []string
	Has(string) bool
	Validate() error
}

type IValidate interface {
	Validate() error
}

var validate *validator.Validate

var errorMessages validator.ValidationErrors

func init() {
	validate = validator.New()
}

func (r *Request) Validator() *validator.Validate {
	return validate
}

func (r *Request) Validate() error {
	return nil
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

func Redirect(path string) {
	routeOfPath := func() *route {
		for _, route := range GetRoutes() {
			if route.Path == path {
				return route
			}
		}
		return &route{}
	}()

	if routeOfPath.Path != "" {
		GetApplication().router.currentRoute = routeOfPath
		RunRouter(routeOfPath.Path)
	}
}

func Back() {
	if referer := GetRequest().request.Referer(); referer != "" {
		u, _ := url.Parse(referer)
		http.Redirect(GetResponse().rw, GetRequest().request, u.Path, http.StatusSeeOther)
	}
	//GetRequest().request.Header.Set("Error-Fields", errorMessages.Error())

	//fmt.Println(GetResponse().rw.Header().Get("Referer"))
	//http.Redirect(GetResponse().rw, GetRequest().request, GetResponse().Header().Get("Referer"), 302)
}

func GetRequest() *Request {
	return GetApplication().req
}
