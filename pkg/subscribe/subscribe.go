package subscribe

import (
	"context"
	"fmt"
	"log"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
	lift "github.com/liftbridge-io/go-liftbridge/liftbridge-grpc"
)

type LBClientConfig struct {
	servers    []string
	streamInfo liftbridge.StreamInfo
	client     liftbridge.Client
}

type ClientIface interface {
	New(servers []string, streamInfo liftbridge.StreamInfo) error
	Sub() error
}

func (lbc *LBClientConfig) New(servers []string, streamInfo liftbridge.StreamInfo) error {
	lbc.streamInfo = streamInfo
	lbc.servers = servers

	client, err := liftbridge.Connect(lbc.servers)
	if err != nil {
		return err
	}
	lbc.client = client
	return nil
}

// func NewLBClientWithConfig(servers []string, streamInfo liftbridge.StreamInfo) (liftbridge.Client, error) {
// 	lbc := LBClientConfig{
// 		servers:    servers,
// 		streamInfo: streamInfo,
// 	}

// 	client, err := liftbridge.Connect(lbc.servers)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }

func (lbClient LBClientConfig) Sub() {
	// Create Liftbridge client.
	servers := []string{"localhost:9292", "localhost:9293", "localhost:9294"}
	stream := liftbridge.StreamInfo{
		Subject:           "foo",
		Name:              "foo-stream",
		ReplicationFactor: 3,
	}

	err := lbClient.New(servers, stream)
	if err != nil {
		log.Panic(err)
	}

	defer lbClient.client.Close()

	// client, err := NewLBClientWithConfig(servers, stream)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer client.Close()

	// Create a stream attached to the NATS subject "foo".
	if err := lbClient.client.CreateStream(context.Background(), stream); err != nil {
		if err != liftbridge.ErrStreamExists {
			panic(err)
		}
	}

	// Subscribe to the stream.
	ctx := context.Background()
	if err := lbClient.client.Subscribe(ctx, stream.Subject, stream.Name, func(msg *lift.Message, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Offset, string(msg.Value))
	}); err != nil {
		panic(err)
	}

	<-ctx.Done()
}
