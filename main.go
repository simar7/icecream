package main

import (
	"context"
	"sync"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
)

type EventStreamClient interface {
	Pub()
	Sub()
	CreateStream(context.Context, liftbridge.StreamInfo) error
}

type LiftBridgeClient struct {
	wg         *sync.WaitGroup
	servers    []string
	streamInfo liftbridge.StreamInfo
	client     liftbridge.Client
}

func main() {
	lbc := LiftBridgeClient{
		wg:      &sync.WaitGroup{},
		servers: []string{"localhost:9292", "localhost:9293", "localhost:9294"}, // TODO(simar7): Add cmd line flag option
		streamInfo: liftbridge.StreamInfo{
			Subject:           "foo",
			Name:              "foo-stream",
			ReplicationFactor: 3,
		},
	}
	lbc.wg.Add(2)
	go sub(lbc)
	go pub(lbc)
	lbc.wg.Wait()
}
