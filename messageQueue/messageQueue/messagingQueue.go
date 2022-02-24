package leafMQ

type MessagingQueue interface {
	Publisher
	Consumer
	Publisher() Publisher
	Consumer() Consumer
}
