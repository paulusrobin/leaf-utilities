package leafServer

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	leafHttpMiddleware "github.com/paulusrobin/leaf-utilities/appRunner/middleware/http"
	"os"
	"strings"
	"time"
)

type (
	HttpServer struct {
		ec                          *echo.Echo
		serviceName, serviceVersion string
		option                      httpOptions
		serverTemplate
	}
)

func NewHttp(serviceName, serviceVersion string, opts ...HttpOption) *HttpServer {
	o := defaultHttpOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &HttpServer{
		ec:             echo.New(),
		serviceName:    serviceName,
		serviceVersion: serviceVersion,
		option:         o,
	}
}

func (s *HttpServer) Serve(sig chan os.Signal) {
	if s.option.validator != nil {
		s.ec.Validator = s.option.validator
	}

	s.ec.Use(
		middleware.Recover(),
		middleware.Gzip(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		}))

	s.ec.GET("/healthz", func(ctx echo.Context) error {
		httpStatus, dependencies := s.option.healthCheck(ctx.Request().Context())
		return ctx.JSON(httpStatus, map[string]interface{}{
			"service_name":    s.serviceName,
			"service_version": s.serviceVersion,
			"status":          "UP",
			"dependencies":    dependencies,
		})
	}, leafHttpMiddleware.Tracer(), leafHttpMiddleware.AccessKey(s.option.healthCheckAccessKey))

	if s.option.errorHandler != nil {
		s.ec.HTTPErrorHandler = func(err error, c echo.Context) {
			s.option.errorHandler(err, c)
		}
	}

	s.serve(sig, serveParam{
		serve: func(sig chan os.Signal) {
			s.option.logger.StandardLogger().Info("[HTTP-SERVER] starting server")
			go func() {
				if err := s.ec.Start(fmt.Sprintf(":%d", s.option.port)); err != nil {
					s.option.logger.StandardLogger().Errorf("[HTTP-SERVER] server interrupted %s", err.Error())
					sig <- os.Interrupt
				}
			}()
			time.Sleep(time.Second)
		},
		register: func() {
			if !s.option.enable {
				return
			}

			if s.option.register != nil {
				s.option.logger.StandardLogger().Debug("[HTTP-SERVER] starting register hooks")
				s.option.register(s.ec, s.option.logger)
			}

			if s.option.featureFlags != nil {
				for _, route := range s.ec.Routes() {
					key := strings.ToLower(route.Method + "_" + route.Path)
					if strings.Contains(route.Path, "healthz") || strings.Contains(route.Path, "docs") {
						continue
					}

					if s.option.featureFlags[key] == nil || s.option.featureFlags[key] == "" {
						s.option.logger.StandardLogger().Errorf("[HTTP-SERVER] shutting down server due to missing %s feature flag", key)
						sig <- os.Interrupt
					}
				}
			}
		},
		beforeRun: func() {
			if s.option.beforeRun != nil {
				s.option.logger.StandardLogger().Debug("[HTTP-SERVER] starting before run hooks")
				s.option.beforeRun(s.ec, s.option.logger)
			}
		},
		afterRun: func() {
			if s.option.afterRun != nil {
				s.option.logger.StandardLogger().Debug("[HTTP-SERVER] starting after run hooks")
				s.option.afterRun(s.ec, s.option.logger)
			}
		},
	})
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), s.option.gracefulPeriod)
	defer cancel()

	s.shutdown(shutdownParam{
		shutdown: func() {
			s.option.logger.StandardLogger().Info("[HTTP-SERVER] shutting down server")
			if err := s.ec.Shutdown(ctx); err != nil {
				s.option.logger.StandardLogger().Errorf("[HTTP-SERVER] server can not be shutdown %s", err.Error())
			}
		},
		beforeExit: func() {
			if s.option.beforeExit != nil {
				s.option.logger.StandardLogger().Debug("[HTTP-SERVER] starting before exit hooks")
				s.option.beforeExit(s.ec, s.option.logger)
			}
		},
		afterExit: func() {
			if s.option.afterExit != nil {
				s.option.logger.StandardLogger().Debug("[HTTP-SERVER] starting after exit hooks")
				s.option.afterExit(s.ec, s.option.logger)
			}
		},
	})
}
