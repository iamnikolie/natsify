package natsify

import (
	"errors"

	"github.com/iamnikolie/natsify/bus"
)

type Meta struct {
	Version  byte
	Encoding bus.Encoding
}

func (h *Meta) Bytes() []byte {
	return []byte{22, h.Version, byte(h.Encoding)}
}

func (h *Meta) Parse(data []byte) ([]byte, error) {
	if len(data) < 3 {
		return data, errors.New("too short")
	}

	if data[0] != 22 {
		return data, errors.New("data is not bus message")
	}

	h.Version = data[1]

	h.Encoding = bus.Encoding(data[2])
	if h.Version != 1 {
		return data, errors.New("bad bus protocol version")
	}

	return data[3:], nil
}

func (h *Meta) Size() int {
	return 3
}
