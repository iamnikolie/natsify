package natsify

import "github.com/iamnikolie/natsify/bus"

func Register(top string, handler func(req *Request)) error {
	return bus.Sub(top, func(msg *bus.Msg) {
		handler(&Request{
			id:       msg.Reply,
			encoding: bus.EncodingJSON,
			topic:    msg.Subject,
			headers:  map[string]interface{}{},
			body:     msg.Data,
		})
	})
}
