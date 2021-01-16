package bus

import (
	"time"

	"github.com/iamnikolie/envigo"
)

// Request ...
func Request(subj string, msg []byte) (*Msg, error) {
	if natsInst == nil || !natsInst.IsConnected() {
		return nil, ErrConnection
	}

	// for default, timeout value is 60 seconds
	t := time.Duration(
		envigo.GetIntFromEnv(
			"NATSIFY_BUS_RPC_TIMEOUT",
			60,
		),
	)

	res, err := natsInst.Request(subj, msg, t*time.Second)
	if err != nil {
		return nil, err
	}
	return &Msg{*res}, nil
}
