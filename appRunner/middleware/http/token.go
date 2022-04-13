package leafHttpMiddleware

import (
	"fmt"
	leafHttpResponse "github.com/enricodg/leaf-utilities/appRunner/response/http"
	"github.com/labstack/echo/v4"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
	"net/http"
)

func Token(token leafToken.Decoder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			var mandatoryBuilder leafMandatory.MandatoryBuilder
			mandatory := leafMandatory.FromContext(eCtx.Request().Context())
			if mandatory.TraceID() == "" {
				builder, err := leafMandatory.NewMandatoryBuilder()
				if err != nil {
					return eCtx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				}
				mandatoryBuilder = httpMandatoryBuilder(eCtx, builder)
				mandatory = mandatoryBuilder.Build()
			} else {
				mandatoryBuilder.Mandatory = mandatory
			}

			if mandatory.Authorization().Authorization() == "" || mandatory.Authorization().Token() == "" {
				return leafHttpResponse.NewJSONResponse(eCtx, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			}

			claims, err := token.DecodeToken(mandatory.Authorization().Token())
			if err != nil {
				return leafHttpResponse.NewJSONResponse(eCtx, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			}

			userLogin := claims.User()
			eCtx.SetRequest(eCtx.Request().WithContext(leafMandatory.Context(eCtx.Request().Context(), mandatoryBuilder.WithUserPhone(userLogin.ID, userLogin.Email, userLogin.Phone).Build())))
			return next(eCtx)
		}
	}
}
