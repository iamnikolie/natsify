package natsify

import (
	"log"
	"net/http"

	"github.com/iamnikolie/natsify/bus"
)

type Request struct {
	id       string
	encoding bus.Encoding
	topic    string
	headers  map[string]interface{}
	body     []byte
}

func (r *Request) Topic() string {
	return r.topic
}

func (r *Request) Unmarshal(i interface{}) error {
	err := decode(r.body, i)
	if err != nil {
		log.Printf("error on decoding message: %s", string(r.body))
	}
	return err
}

func (r *Request) Header(name string, v interface{}) *Request {
	r.headers[name] = v
	return r
}

func (r *Request) Respond(data interface{}, code ...interface{}) error {
	res := Response{
		Code:    http.StatusOK,
		Headers: r.headers,
		Data:    data,
	}

	if len(code) > 0 {
		res.Code = code[0].(int)
		if len(code) == 2 {
			res.Bytes = code[1].([]byte)
		}
	}

	msg, err := encode(res, r.encoding)
	if err != nil {
		return err
	}
	return bus.Pub(r.id, msg)
}

func (r *Request) Success() error {
	return r.Respond(map[string]interface{}{
		"success": true,
	}, http.StatusOK)
}

func (r *Request) Error(desc string, code int) error {
	return r.Respond(
		map[string]interface{}{
			"error":   true,
			"message": desc,
		},
		code,
	)
}
