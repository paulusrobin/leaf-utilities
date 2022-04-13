package leafMiddleware

import "github.com/labstack/echo/v4"

func GetHttpRequestHeader(eCtx echo.Context, headers ...string) string {
	var value = ""
	for _, header := range headers {
		value = eCtx.Request().Header.Get(header)
		if value != "" {
			break
		}
	}
	return value
}

func GetHttpRequestHeaderWithDefault(eCtx echo.Context, fn func() string, headers ...string) string {
	var value = ""
	for _, header := range headers {
		value = eCtx.Request().Header.Get(header)
		if value != "" {
			break
		}
	}
	if value == "" {
		return fn()
	}
	return value
}
