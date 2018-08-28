package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/liftbridge-io/go-liftbridge"
	"github.com/stretchr/testify/assert"
)

func (mlbc MockLiftBridgeClient) CreateStream(ctx context.Context, stream liftbridge.StreamInfo) error {
	fmt.Println("fake CreateStream() called")
	return nil
}

func Test_Client_CreateStream(t *testing.T) {
	mlbc := MockLiftBridgeClient{
		streamInfo: liftbridge.StreamInfo{
			Subject:           "foo",
			Name:              "foo-stream",
			ReplicationFactor: 3,
		},
	}
	assert.NoError(t, createStream(mlbc, context.Background(), mlbc.streamInfo))
}
