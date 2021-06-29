package filter

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-logger"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(TraceFilter))
}

const componentIDGOHttpServer = 5004

type TraceFilter struct {
	_              SpringWeb.Filter `export:""`
	Enabled        bool             `value:"${trace.enabled}"`
	BackendAddress string           `value:"${trace.backend-address}"`

	SpringApplicationName string `value:"${spring.application.name}"`

	tracer   *go2sky.Tracer
	initOnce sync.Once
}

func (filter *TraceFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	if filter.Enabled == false {
		chain.Next(webCtx)
		return
	}

	// init go2sky tracer
	filter.initOnce.Do(func() {
		r, err := reporter.NewGRPCReporter(filter.BackendAddress)
		if err != nil {
			SpringLogger.Errorf("new gRPC reporter err: %v")
			return
		}
		tracer, err := go2sky.NewTracer(filter.SpringApplicationName, go2sky.WithReporter(r))
		if err != nil {
			SpringLogger.Errorf("new tracer err:: %v")
			return
		}
		filter.tracer = tracer
	})

	if filter.tracer == nil {
		chain.Next(webCtx)
		return
	}

	// trace begin, create entry span
	request := webCtx.Request()
	span, _, err := filter.tracer.CreateEntrySpan(webCtx.Context(), fmt.Sprintf("/%s%s", request.Method, request.URL.Path), func(headerKey string) (string, error) {
		return webCtx.GetHeader(propagation.Header), nil
	})
	if err != nil {
		SpringLogger.Debugf("create entry span err: %v", err)
		chain.Next(webCtx)
		return
	}

	span.SetComponent(componentIDGOHttpServer)
	span.Tag(go2sky.TagHTTPMethod, request.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", request.Host, request.URL.Path))
	defer func() {
		nativeContext := webCtx.NativeContext()
		c := nativeContext.(*gin.Context)

		if len(c.Errors) > 0 {
			span.Error(time.Now(), c.Errors.String())
		}
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(c.Writer.Status()))
		span.End()
	}()
	chain.Next(webCtx)
}
