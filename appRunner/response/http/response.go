package leafHttpResponse

import "github.com/labstack/echo/v4"

const ApiResponse = "api-response"

type Response interface {
	Val() interface{}
}

func NewJSONResponse(
	eCtx echo.Context,
	statusCode int,
	data interface{},
) error {
	var response Response

	if data == nil {
		return eCtx.JSON(statusCode, nil)
	}

	if statusCode >= 400 {
		response = newErrorResponse(eCtx.Request().Context(), statusCode, data.(error))
	} else {
		response = newSuccessResponse(data)
	}

	eCtx.Set(ApiResponse, response.Val())
	return eCtx.JSON(statusCode, response.Val())
}
