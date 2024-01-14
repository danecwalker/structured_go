package event

// Event is the interface that all events must implement.
type Event interface {
	// Name returns the name of the event.
	Name() string
}

// EventHandler is the interface that all event handlers must implement.
type EventHandler interface {
	// Handle handles an event.
	Handle(event Event) error
}
