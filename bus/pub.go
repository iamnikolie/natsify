package bus

import "log"

func Pub(top string, msg []byte) error {
	if natsInst == nil || !natsInst.IsConnected() {
		return ErrConnection
	}

	err := natsInst.Publish(top, msg)
	if err != nil {
		log.Printf("error while publishing: %s", err.Error())
	}

	natsInst.Flush()
	return nil
}
