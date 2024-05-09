package pubsub

// Publisher is the interface that defines the methods that a publisher must implement in CQRS pattern.

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Publisher interface {
	Publish(ctx context.Context, e endpoint.Endpoint, request interface{}) error
}
