package leafHttpMiddleware

import (
	"bytes"
	"fmt"
	leafHttpResponse "github.com/enricodg/leaf-utilities/appRunner/response/http"
	"github.com/labstack/echo/v4"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"github.com/pkg/errors"
	"io/ioutil"
)

func Tracer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			operationName := fmt.Sprintf("HTTP - [%s] %s", eCtx.Request().Method, eCtx.Path())
			span, context := tracer.StartSpanFromContext(
				eCtx.Request().Context(), operationName,
				tracer.Tag("http-request", eCtx.Request()),
				tracer.Tag("http-response-writer", eCtx.Response().Writer),
			)

			// Request ID
			mandatory := leafMandatory.FromContext(eCtx.Request().Context())
			span.SetTag("requestId", mandatory.TraceID())
			span.SetBaggageItem("requestId", mandatory.TraceID())

			// Request Header
			var (
				headerRequestBytes []byte = []byte("")
				mapRequestHeader   map[string]interface{}
			)
			headerRequestBytes, _ = json.Marshal(eCtx.Request().Header)
			_ = json.Unmarshal(headerRequestBytes, &mapRequestHeader)
			span.SetTag("requestHeader", eCtx.Request().Header)
			span.SetBaggageItem("requestHeader", string(headerRequestBytes))

			// Request Body
			if eCtx.Request().Body != nil {
				bodyRequestBytes, _ := ioutil.ReadAll(eCtx.Request().Body)
				eCtx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyRequestBytes)) // Reset
				span.SetTag("requestBody", string(bodyRequestBytes))
				span.SetBaggageItem("requestBody", string(bodyRequestBytes))
			}

			// Request Path
			requestPath := fmt.Sprintf("[%s] %s", eCtx.Request().Method, eCtx.Path())
			span.SetTag("requestPath", requestPath)
			span.SetBaggageItem("requestPath", requestPath)
			span.SetTag("requestRawPath", eCtx.Request().URL.String())
			span.SetBaggageItem("requestRawPath", eCtx.Request().URL.String())

			eCtx.SetRequest(eCtx.Request().WithContext(context))
			err := next(eCtx)

			// Response Body
			var mapResponse map[string]interface{}
			if eCtx.Get(leafHttpResponse.ApiResponse) != nil {
				var bodyResponseBytes []byte = []byte("")
				bodyResponseBytes, _ = json.Marshal(eCtx.Get(leafHttpResponse.ApiResponse))
				_ = json.Unmarshal(bodyResponseBytes, &mapResponse)
				span.SetTag("responseBody", mapResponse)
				span.SetBaggageItem("responseBody", string(bodyResponseBytes))
			}

			// Response Header
			if nil != eCtx.Response() {
				var headerResponseBytes []byte = []byte("")
				headerResponseBytes, _ = json.Marshal(eCtx.Response().Header())
				_ = json.Unmarshal(headerResponseBytes, &mapResponse)
				span.SetTag("responseHeader", eCtx.Response().Header())
				span.SetBaggageItem("responseHeader", string(headerResponseBytes))

				span.SetTag("responseStatus", eCtx.Response().Status)
				span.SetBaggageItem("responseStatus", fmt.Sprintf("%d", eCtx.Response().Status))
			}

			if nil != err {
				if !leafFunctions.SkipNoticeError(eCtx.Request().Context()) {
					span.Finish(tracer.WithError(errors.Wrap(err, requestPath)))
				}
				return err
			}

			span.Finish()
			return nil
		}
	}
}
