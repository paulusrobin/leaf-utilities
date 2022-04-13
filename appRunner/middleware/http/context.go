package leafHttpMiddleware

import (
	"bytes"
	"fmt"
	leafMiddleware "github.com/enricodg/leaf-utilities/appRunner/middleware"
	leafHttpResponse "github.com/enricodg/leaf-utilities/appRunner/response/http"
	"github.com/labstack/echo/v4"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func AppContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return eCtx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}

			mandatory := httpMandatoryBuilder(eCtx, builder).Build()
			ctx := leafMandatory.Context(eCtx.Request().Context(), mandatory)
			eCtx.SetRequest(eCtx.Request().WithContext(ctx))
			return next(eCtx)
		}
	}
}

func AppContextWithLogger(logger leafLogger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return eCtx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			}

			mandatory := httpMandatoryBuilder(eCtx, builder).Build()
			ctx := leafMandatory.Context(eCtx.Request().Context(), mandatory)
			eCtx.SetRequest(eCtx.Request().WithContext(ctx))
			logRequest(eCtx, logger)

			start := leafTime.Now()
			err = next(eCtx)
			logResponse(eCtx, logger, leafTime.Now().Sub(start), err)
			return err
		}
	}
}

func logRequest(eCtx echo.Context, logger leafLogger.Logger) {
	var mapRequest map[string]interface{}
	if eCtx.Request().Body != nil {
		requestByte, _ := ioutil.ReadAll(eCtx.Request().Body)
		eCtx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(requestByte)) // Reset
		_ = json.Unmarshal(requestByte, &mapRequest)
	}

	requestMsg := fmt.Sprintf("[api-log] - [REQUEST] - [%s] %s",
		eCtx.Request().Method,
		eCtx.Request().URL.Path)

	if eCtx.Request().URL.RawQuery != "" {
		requestMsg += fmt.Sprintf("?%s", eCtx.Request().URL.RawQuery)
	}

	logger.Info(leafLogger.BuildMessage(eCtx.Request().Context(), requestMsg,
		leafLogger.WithAttr("payload", mapRequest),
		leafLogger.WithAttr("headers", eCtx.Request().Header),
	))
}

func logResponse(eCtx echo.Context, logger leafLogger.Logger, duration time.Duration, err error) {
	responseMsg := fmt.Sprintf("[api-log] - [RESPONSE] - [%d] [%dms] [%s] %s",
		eCtx.Response().Status,
		duration.Milliseconds(),
		eCtx.Request().Method,
		eCtx.Request().URL.Path)
	if eCtx.Request().URL.RawQuery != "" {
		responseMsg += fmt.Sprintf("?%s", eCtx.Request().URL.RawQuery)
	}

	if err != nil {
		logger.Error(leafLogger.BuildMessage(eCtx.Request().Context(), responseMsg,
			leafLogger.WithAttr("error", err)))
	} else {
		logger.Info(leafLogger.BuildMessage(eCtx.Request().Context(), responseMsg,
			leafLogger.WithAttr("response", eCtx.Get(leafHttpResponse.ApiResponse))))
	}
}

func httpMandatoryBuilder(eCtx echo.Context, builder leafMandatory.MandatoryBuilder) leafMandatory.MandatoryBuilder {
	mandatoryBuilder := builder.
		WithTraceID(leafMiddleware.GetHttpRequestHeaderWithDefault(eCtx, leafMiddleware.GenerateUUID, leafHeader.TraceID, leafHeader.XTraceID)).
		WithDeviceType(leafMiddleware.GetHttpRequestHeaderWithDefault(eCtx, leafMiddleware.EmptyString, leafHeader.DeviceType)).
		WithAuthorization(leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.Authorization)).
		WithLanguage(leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.Language)).
		WithApiKey(leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.ApiKey, leafHeader.XApiKey)).
		WithIpAddresses(strings.Split(leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.IpAddress), ",")).
		WithUserAgent(leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.UserAgent)).
		WithServiceSecret(
			leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.ServiceID, leafHeader.XServiceID),
			leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.ServiceSecret, leafHeader.XServiceSecret),
		)

	appsVersion := leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.AppVersion, leafHeader.XAppVersion)
	deviceID := leafMiddleware.GetHttpRequestHeader(eCtx, leafHeader.DeviceID)
	if deviceID != "" || appsVersion != "" {
		mandatoryBuilder = mandatoryBuilder.WithApplication(deviceID, appsVersion)
	}
	return mandatoryBuilder
}
