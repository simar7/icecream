package main

import (
	"fmt"
	"sync"
)

type EventStreamClient interface {
	Pub()
	Sub()
}

type LiftBridgeClient struct {
	wg *sync.WaitGroup
}

func (lbc LiftBridgeClient) Pub() {
	defer lbc.wg.Done()
	fmt.Println("real Pub() called")
}

func pub(lbcIface EventStreamClient) {
	lbcIface.Pub()
}

func (lbc LiftBridgeClient) Sub() {
	defer lbc.wg.Done()
	fmt.Println("real Sub() called")
}

func sub(lbcIface EventStreamClient) {
	lbcIface.Sub()
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
