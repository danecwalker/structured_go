package event

type Broker interface {
	// Publish publishes an event to the broker.
	Publish(event Event) error

	// Subscribe subscribes to an event on the broker.
	Subscribe(event Event, handler EventHandler) error

	// Unsubscribe unsubscribes from an event on the broker.
	Unsubscribe(event Event, handler EventHandler) error

	// Close closes the broker.
	Close() error
}
