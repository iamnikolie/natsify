package natsify

import (
	"encoding/json"

	"net/http"

	"github.com/iamnikolie/natsify/bus"
)

func Call(
	topic string,
	args *map[string]interface{},
	res *Response,
	data ...interface{},
) error {
	if args == nil {
		e := make(map[string]interface{})
		args = &e
	}

	bt, err := encode(args, detectEncoding(args))
	if err != nil {
		return err
	}

	msg, err := bus.Request(topic, bt)
	if err != nil {
		return err
	}

	raw := map[string]json.RawMessage{}
	if err := decode(msg.Data, &raw); err != nil {
		return err
	}

	if err := json.Unmarshal(raw["code"], &res.Code); err != nil {
		return err
	}

	if v, ok := raw["headers"]; ok {
		if err := json.Unmarshal(v, &res.Headers); err != nil {
			return err
		}
	}

	res.Bytes = raw["data"]

	if res.Code >= http.StatusBadRequest {
		res.Data = string(raw["data"])
		return nil
	}

	if v, ok := raw["data"]; ok {

		if len(data) > 0 && data[0] != nil {
			if err := json.Unmarshal(v, data[0]); err != nil {
				return err
			}
			res.Data = data[0]
		} else {
			if err := json.Unmarshal(v, &res.Data); err != nil {
				return err
			}
		}
	}

	return nil
}
