package main

import (
	"fmt"
	"sync"
	"testing"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
	"github.com/stretchr/testify/assert"
)

type MockLiftBridgeClient struct {
	wg         *sync.WaitGroup
	streamInfo liftbridge.StreamInfo
}

func (mlbc MockLiftBridgeClient) Pub() {
	defer mlbc.wg.Done()
	fmt.Println("fake Pub() called")
}

func (mlbc MockLiftBridgeClient) Sub() {
	defer mlbc.wg.Done()
	fmt.Println("fake Sub() called")
}

func TestPub(t *testing.T) {
	f := MockLiftBridgeClient{
		wg: &sync.WaitGroup{},
	}
	f.wg.Add(1)
	go pub(f)
	assert.NotPanics(t, f.wg.Wait)
}

func TestSub(t *testing.T) {
	f := MockLiftBridgeClient{
		wg: &sync.WaitGroup{},
	}
	f.wg.Add(1)
	go sub(f)
	assert.NotPanics(t, f.wg.Wait)

}

func TestPubAndSub(t *testing.T) {
	f := MockLiftBridgeClient{
		wg: &sync.WaitGroup{},
	}

	f.wg.Add(2)
	go pub(f)
	go sub(f)
	assert.NotPanics(t, f.wg.Wait)
}
