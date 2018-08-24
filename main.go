package main

import (
	"sync"
)

type EventStreamClient interface {
	Pub()
	Sub()
}

type LiftBridgeClient struct {
	wg *sync.WaitGroup
}

func main() {
	lbc := LiftBridgeClient{
		wg: &sync.WaitGroup{},
	}
	lbc.wg.Add(2)
	go sub(lbc)
	go pub(lbc)
	lbc.wg.Wait()
}
