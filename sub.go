package main

import (
	"context"
	"fmt"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
	lift "github.com/liftbridge-io/go-liftbridge/liftbridge-grpc"
)

func (lbc LiftBridgeClient) Sub() {
	defer lbc.wg.Done()
	fmt.Println("real Sub() called")

	// TODO(simar7): Add exponential backoff rather than panic
	client, err := liftbridge.Connect(lbc.servers)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.CreateStream(context.Background(), lbc.streamInfo); err != nil {
		if err != liftbridge.ErrStreamExists {
			panic(err)
		}
	}

	ctx := context.Background()
	if err := client.Subscribe(ctx, lbc.streamInfo.Subject, lbc.streamInfo.Name, func(msg *lift.Message, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Offset, string(msg.Value))
	}); err != nil {
		panic(err)
	}

	<-ctx.Done()
}

func sub(lbcIface EventStreamClient) {
	lbcIface.Sub()
}
