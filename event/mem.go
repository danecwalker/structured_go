package event

type MemBroker struct {
	// Map of event names to event handlers.
	handlers map[string][]EventHandler
}

// NewMemBroker returns a new memory broker.
func NewMemBroker() *MemBroker {
	return &MemBroker{
		handlers: make(map[string][]EventHandler),
	}
}

// Publish publishes an event to the broker.
func (b *MemBroker) Publish(event Event) error {
	for _, handler := range b.handlers[event.Name()] {
		if err := handler.Handle(event); err != nil {
			return err
		}
	}
	return nil
}

// Subscribe subscribes to an event on the broker.
func (b *MemBroker) Subscribe(event Event, handler EventHandler) error {
	b.handlers[event.Name()] = append(b.handlers[event.Name()], handler)
	return nil
}

// Unsubscribe unsubscribes from an event on the broker.
func (b *MemBroker) Unsubscribe(event Event, handler EventHandler) error {
	handlers := b.handlers[event.Name()]
	for i, h := range handlers {
		if h == handler {
			b.handlers[event.Name()] = append(handlers[:i], handlers[i+1:]...)
			return nil
		}
	}
	return nil
}

// Close closes the broker.
func (b *MemBroker) Close() error {
	b.handlers = make(map[string][]EventHandler)
	return nil
}
