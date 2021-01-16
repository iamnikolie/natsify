package bus

import (
	"errors"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	// ReconnectionTimeout ...
	ReconnectionTimeout = 1 * time.Minute

	EncodingBinary Encoding = 0
	EncodingJSON   Encoding = 1
	EncodingBSON   Encoding = 3

	EncodingKey = "encv1"
)

var (
	// ErrConnection ...
	ErrConnection = errors.New("not found")

	natsRegistryName string
	natsInst         *nats.Conn
)

// Ready ...
func Ready() bool {
	return !(natsInst == nil || !natsInst.IsConnected())
}
