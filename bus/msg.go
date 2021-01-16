package bus

import "github.com/nats-io/nats.go"

type Msg struct {
	nats.Msg
}
