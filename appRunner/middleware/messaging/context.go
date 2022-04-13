package leafMessagingMiddleware

import (
	"context"
	"fmt"
	leafMiddleware "github.com/enricodg/leaf-utilities/appRunner/middleware"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"strings"
	"time"
)

func AppContext() leafMQ.MiddlewareFunc {
	return func(next leafMQ.MiddlewareHandlerFunc) leafMQ.MiddlewareHandlerFunc {
		return func(ctx context.Context, msg leafMQ.DispatchDTO) error {
			if ctx == nil {
				ctx = context.Background()
			}

			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return err
			}

			ctx = leafMandatory.Context(ctx, messagingMandatoryBuilder(msg, builder).Build())
			return next(ctx, msg)
		}
	}
}

func AppContextWithLogger(logger leafLogger.Logger) leafMQ.MiddlewareFunc {
	return func(next leafMQ.MiddlewareHandlerFunc) leafMQ.MiddlewareHandlerFunc {
		return func(ctx context.Context, msg leafMQ.DispatchDTO) error {
			if ctx == nil {
				ctx = context.Background()
			}

			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return err
			}

			mandatory := messagingMandatoryBuilder(msg, builder).Build()
			ctx = leafMandatory.Context(ctx, mandatory)
			logRequest(ctx, msg, logger)

			start := leafTime.Now()
			err = next(ctx, msg)
			duration := leafTime.Now().Sub(start)

			logResponse(ctx, msg, logger, duration, err)
			return err
		}
	}
}

func logRequest(ctx context.Context, msg leafMQ.DispatchDTO, logger leafLogger.Logger) {
	requestMsg := fmt.Sprintf(
		"[messaging-log] - [REQUEST] - listening topic %s", msg.Source)

	logger.Info(leafLogger.BuildMessage(ctx, requestMsg,
		leafLogger.WithAttr("payload", string(msg.Msg.Data)),
		leafLogger.WithAttr("headers", msg.Msg.Attributes),
		leafLogger.WithAttr("msgType", msg.MsgType),
	))
}

func logResponse(ctx context.Context, msg leafMQ.DispatchDTO, logger leafLogger.Logger, duration time.Duration, err error) {
	responseMsg := fmt.Sprintf("[messaging-log] - [RESPONSE] - [%dms] listening topic %s", duration.Milliseconds(), msg.Source)

	if err != nil {
		logger.Error(leafLogger.BuildMessage(ctx, responseMsg,
			leafLogger.WithAttr("payload", string(msg.Msg.Data)),
			leafLogger.WithAttr("headers", msg.Msg.Attributes),
			leafLogger.WithAttr("msgType", msg.MsgType),
			leafLogger.WithAttr("error", err),
		))
	} else {
		logger.Info(leafLogger.BuildMessage(ctx, responseMsg,
			leafLogger.WithAttr("payload", string(msg.Msg.Data)),
			leafLogger.WithAttr("headers", msg.Msg.Attributes),
			leafLogger.WithAttr("msgType", msg.MsgType),
		))
	}
}

func messagingMandatoryBuilder(msg leafMQ.DispatchDTO, builder leafMandatory.MandatoryBuilder) leafMandatory.MandatoryBuilder {
	mandatoryBuilder := builder.
		WithTraceID(leafMiddleware.GetMessagingRequestHeaderWithDefault(msg.Msg, leafMiddleware.GenerateUUID, leafHeader.MessagingTraceID)).
		WithDeviceType(leafMiddleware.GetMessagingRequestHeaderWithDefault(msg.Msg, leafMiddleware.EmptyString, leafHeader.MessagingDeviceType)).
		WithAuthorization(leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingAuthorization)).
		WithLanguage(leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingLang)).
		WithApiKey(leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingApiKey)).
		WithIpAddresses(strings.Split(leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingIpAddress), ",")).
		WithUserAgent(leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingUserAgent)).
		WithServiceSecret(
			leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingServiceID),
			leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingServiceSecret),
		)

	appsVersion := leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingAppVersion)
	deviceID := leafMiddleware.GetMessagingRequestHeader(msg.Msg, leafHeader.MessagingDeviceID)
	if deviceID != "" || appsVersion != "" {
		mandatoryBuilder = mandatoryBuilder.WithApplication(deviceID, appsVersion)
	}
	return mandatoryBuilder
}
