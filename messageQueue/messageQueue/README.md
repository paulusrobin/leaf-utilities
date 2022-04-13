# Message Queue

## Message Queue Interface
```go
MessageQueue interface{
    Publish(ctx context.Context, topic string, msg Message) error	
    Use(middlewareFunc ...MiddlewareFunc)
    Listen()
    Subscribe(topic string, dispatcher Dispatcher) error
    Ping(ctx context.Context) error
    Close() error  
    Publisher() Publisher
    Consumer() Consumer
}
```