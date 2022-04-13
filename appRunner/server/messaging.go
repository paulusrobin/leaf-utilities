package leafServer

import (
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"os"
	"time"
)

type (
	MessagingServer struct {
		consumer leafMQ.Consumer
		option   consumerOptions
		serverTemplate
	}
)

func NewMessaging(consumer leafMQ.Consumer, opts ...ConsumerOption) *MessagingServer {
	o := defaultMessagingOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	return &MessagingServer{
		consumer: consumer,
		option:   o,
	}
}

func (s *MessagingServer) Serve(sig chan os.Signal) {
	if !s.option.enable {
		s.option.logger.StandardLogger().Info("[MESSAGING-SERVER] server disabled")
		return
	}

	s.serve(sig, serveParam{
		serve: func(sig chan os.Signal) {
			s.option.logger.StandardLogger().Info("[MESSAGING-SERVER] starting server")
			go s.consumer.Listen()
			time.Sleep(time.Second)
		},
		register: func() {
			if s.option.register != nil {
				s.option.logger.StandardLogger().Debug("[MESSAGING-SERVER] starting register hooks")
				s.option.register(s.consumer, s.option.logger)
			}
		},
		beforeRun: func() {
			if s.option.beforeRun != nil {
				s.option.logger.StandardLogger().Debug("[MESSAGING-SERVER] starting before run hooks")
				s.option.beforeRun(s.consumer, s.option.logger)
			}
		},
		afterRun: func() {
			if s.option.afterRun != nil {
				s.option.logger.StandardLogger().Debug("[MESSAGING-SERVER] starting after run hooks")
				s.option.afterRun(s.consumer, s.option.logger)
			}
		},
	})
}

func (s *MessagingServer) Shutdown() {
	if !s.option.enable {
		return
	}

	s.shutdown(shutdownParam{
		shutdown: func() {
			s.option.logger.StandardLogger().Info("[MESSAGING-SERVER] shutting down server")
			if err := s.consumer.Close(); err != nil {
				s.option.logger.StandardLogger().Errorf("[MESSAGING-SERVER] server can not be shutdown %s", err.Error())
			}
		},
		beforeExit: func() {
			if s.option.beforeExit != nil {
				s.option.logger.StandardLogger().Debug("[MESSAGING-SERVER] starting before exit hooks")
				s.option.beforeExit(s.consumer, s.option.logger)
			}
		},
		afterExit: func() {
			if s.option.afterExit != nil {
				s.option.logger.StandardLogger().Debug("[MESSAGING-SERVER] starting after exit hooks")
				s.option.afterExit(s.consumer, s.option.logger)
			}
		},
	})
}
