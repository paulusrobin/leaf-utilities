package leafHttpMiddleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const accessKey = `Access-Key`

func AccessKey(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			if "" == key {
				return next(eCtx)
			}

			if key != eCtx.Request().Header.Get(accessKey) {
				return eCtx.JSON(http.StatusForbidden, map[string]interface{}{
					"error":   "forbidden",
					"message": "endpoint cannot be accessed",
				})
			}
			return next(eCtx)
		}
	}
}
