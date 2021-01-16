package natsify

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/iamnikolie/natsify/bus"
)

func detectEncoding(args *map[string]interface{}) bus.Encoding {
	out := bus.EncodingJSON
	mp := *args
	if raw, ok := mp[bus.EncodingKey]; ok {
		switch raw.(type) {
		case byte:
			out = bus.Encoding(raw.(byte))
			break
		default:
		}
	}
	return out
}

func encode(o interface{}, enc bus.Encoding) (out []byte, err error) {
	mcap := Meta{1, enc}
	switch enc {
	case bus.EncodingBinary:
		v := reflect.ValueOf(o)
		if v.Kind() != reflect.Slice || v.Type() != reflect.TypeOf([]byte(nil)) {
			return out, errors.New("data is not byte array")
		}
		out = o.([]byte)
		break
	case bus.EncodingJSON:
		out, err = json.Marshal(o)
		break
	default:
		err = errors.New("unknown or unsupported encoding in call")
	}
	if err == nil {
		h := mcap.Bytes()
		out = append(h, out...)
	}
	return out, err
}

func decode(msg []byte, o interface{}) error {
	mcap := Meta{}
	data, err := mcap.Parse(msg)
	if err != nil {
		return err
	}
	switch mcap.Encoding {
	case bus.EncodingBinary:
		o = &data
		break
	case bus.EncodingJSON:
		return json.Unmarshal(data, o)
	}
	return errors.New("unknown or unsupported encoding in rpc message")
}
