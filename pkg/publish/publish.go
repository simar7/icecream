package publish

import (
	"time"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
	nats "github.com/nats-io/go-nats"
)

func Pub() {
	conn, err := nats.GetDefaultOptions().Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Flush()
	defer conn.Close()

	// Publish message every 1 second
	for {
		msg := liftbridge.NewMessage([]byte("value"), liftbridge.MessageOptions{Key: []byte("key")})
		if err := conn.Publish("foo", msg); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}

}
