package bus

import (
	"log"

	"github.com/nats-io/nats.go"
)

// Sub ...
func Sub(top string, handler func(msg *Msg)) error {
	var err error
	if Ready() {
		if _, err = natsInst.Subscribe(top, func(nm *nats.Msg) {
			handler(&Msg{*nm})
		}); err != nil {
			log.Printf(
				"error while subscribe on topic %s: %s",
				top,
				err.Error(),
			)
		}
	} else {
		log.Printf("connected to topic: %s", top)
	}
	return err
}
