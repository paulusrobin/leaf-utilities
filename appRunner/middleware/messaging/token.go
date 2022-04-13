package leafMessagingMiddleware

import (
	"context"
	"fmt"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
)

func Token(token leafToken.Decoder) leafMQ.MiddlewareFunc {
	return func(next leafMQ.MiddlewareHandlerFunc) leafMQ.MiddlewareHandlerFunc {
		return func(ctx context.Context, msg leafMQ.DispatchDTO) error {
			if ctx == nil {
				ctx = context.Background()
			}

			var mandatoryBuilder leafMandatory.MandatoryBuilder
			mandatory := leafMandatory.FromContext(ctx)
			if mandatory.TraceID() == "" {
				builder, err := leafMandatory.NewMandatoryBuilder()
				if err != nil {
					return err
				}
				mandatoryBuilder = messagingMandatoryBuilder(msg, builder)
				mandatory = mandatoryBuilder.Build()
			} else {
				mandatoryBuilder.Mandatory = mandatory
			}

			if mandatory.Authorization().Authorization() == "" || mandatory.Authorization().Token() == "" {
				return fmt.Errorf("invalid authorization")
			}

			claims, err := token.DecodeToken(mandatory.Authorization().Token())
			if err != nil {
				return err
			}

			userLogin := claims.User()
			ctx = leafMandatory.Context(ctx, mandatoryBuilder.WithUserPhone(userLogin.ID, userLogin.Email, userLogin.Phone).Build())
			return next(ctx, msg)
		}
	}
}
