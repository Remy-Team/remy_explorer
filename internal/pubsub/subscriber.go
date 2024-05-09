package pubsub

import "github.com/go-kit/kit/endpoint"

// Subscriber is the interface that defines the methods that a subscriber must implement in CQRS pattern.

type Subscriber interface {
	Subscribe(e endpoint.Endpoint) error
}
