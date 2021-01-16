package bus

import (
	"log"
	"time"

	"github.com/iamnikolie/envigo"

	"github.com/nats-io/nats.go"
)

// Do ...
func Do() {
	reconfigLoop()
}

func reconfigLoop() {
	configureOnce()
	go func() {
		for {
			configureOnce()

			time.Sleep(ReconnectionTimeout)
		}
	}()
}

func configureOnce() {
	url := "nats://" + envigo.GetStrFromEnv(
		"NATSIFY_HOST",
		"127.0.0.1:4242",
	)
	opts := nats.GetDefaultOptions()
	opts.MaxReconnect = -1
	opts.ReconnectWait = ReconnectionTimeout
	opts.Url = url
	if natsInst != nil && natsInst.ConnectedUrl() == url {
		return
	}
	opts.DisconnectedCB = func(nc *nats.Conn) {
		log.Printf("message bus disconnected")
	}
	opts.ReconnectedCB = func(nc *nats.Conn) {
		log.Printf("try to connect to NATS on: %s\n", nc.ConnectedUrl())
	}
	nc, err := opts.Connect()
	if err != nil {
		log.Printf("error on connect to %s: %s", url, err.Error())
		return
	}
	if natsInst != nil {
		log.Printf(
			"change message bus connection from %s to %s",
			natsInst.ConnectedUrl(),
			nc.ConnectedUrl(),
		)
		natsInst.Close()
	} else {
		log.Printf(
			"connected to message bus on: %s",
			nc.ConnectedUrl(),
		)
	}
	natsInst = nc
}
