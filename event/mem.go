package event

type MemBroker struct {
	// Map of event names to event handlers.
	handlers map[string][]EventHandler
	queue    chan Event
}

// NewMemBroker returns a new memory broker.
func NewMemBroker() *MemBroker {
	return &MemBroker{
		handlers: make(map[string][]EventHandler),
		queue:    make(chan Event),
	}
}

// Run runs the broker.
func (b *MemBroker) Run() {
	for event := range b.queue {
		for _, handler := range b.handlers[event.Name()] {
			go handler.Handle(event)
		}
	}
}

// Publish publishes an event to the broker.
func (b *MemBroker) Publish(event Event) {
	b.queue <- event
}

// Subscribe subscribes to an event on the broker.
func (b *MemBroker) Subscribe(event Event, handler EventHandler) {
	b.handlers[event.Name()] = append(b.handlers[event.Name()], handler)
}

// Unsubscribe unsubscribes from an event on the broker.
func (b *MemBroker) Unsubscribe(event Event, handler EventHandler) {
	handlers := b.handlers[event.Name()]
	for i, h := range handlers {
		if h == handler {
			b.handlers[event.Name()] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
}

// Close closes the broker.
func (b *MemBroker) Close() error {
	close(b.queue)
	return nil
}
