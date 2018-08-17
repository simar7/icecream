package subscribe

import (
	"context"
	"errors"
	"testing"

	liftbridge "github.com/liftbridge-io/go-liftbridge"
)

type mockSubClient struct {
	new func(servers []string, streamInfo liftbridge.StreamInfo) error
	sub func() error
}

func (msc mockSubClient) New(servers []string, streamInfo liftbridge.StreamInfo) error {
	if msc.new != nil {
		return msc.New(servers, streamInfo)
	}

	return nil
}

func (msc mockSubClient) Sub() error {
	if msc.sub != nil {
		return msc.Sub()
	}
	return nil
}

type mockLiftBridgeClient struct {
	close        func() error
	createStream func(ctx context.Context, stream liftbridge.StreamInfo) error
	subscribe    func(ctx context.Context, subject, name string, handler liftbridge.Handler, opts ...liftbridge.SubscriptionOption) error
}

func (mlbc mockLiftBridgeClient) Close() error {
	if mlbc.close != nil {
		return mlbc.Close()
	}
	return nil
}

func (mlbc mockLiftBridgeClient) CreateStream(ctx context.Context, stream liftbridge.StreamInfo) error {
	if mlbc.createStream != nil {
		return mlbc.CreateStream(ctx, stream)
	}
	return nil
}

func (mlbc mockLiftBridgeClient) Subscribe(ctx context.Context, subject, name string, handler liftbridge.Handler, opts ...liftbridge.SubscriptionOption) error {
	if mlbc.subscribe != nil {
		return mlbc.Subscribe(ctx, subject, name, handler, opts...)
	}
	return nil
}

func TestSub(t *testing.T) {

	// l := LBClientConfig{}
	l := mockSubClient{
		new: func(servers []string, streamInfo liftbridge.StreamInfo) error {
			return errors.New("wtf?")
		},
	}
	l.Sub()

}
