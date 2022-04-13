package leafMQ

type MessageQueue interface {
	Publisher
	Consumer
	Publisher() Publisher
	Consumer() Consumer
}
