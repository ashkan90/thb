package system

type Request struct {
	name string
}

func GetRequest() *Request {
	return GetApplication().req
}